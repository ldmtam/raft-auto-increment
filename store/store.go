package store

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	pb "github.com/ldmtam/raft-auto-increment/auto_increment/pb"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"github.com/ldmtam/raft-auto-increment/config"
	"github.com/ldmtam/raft-auto-increment/database"
	"github.com/ldmtam/raft-auto-increment/database/boltdb"
	"google.golang.org/grpc"
)

const (
	retainSnapshotCount = 2
	logCacheCapacity    = 512
	applyTimeout        = 10 * time.Second
)

// Store represents the Store structure
type Store struct {
	raft *raft.Raft

	db     database.AutoIncrement
	config *config.Config
}

// New return new instance of Store object
func New(config *config.Config) (*Store, error) {
	store := &Store{
		config: config,
	}

	db, err := boltdb.New(store.config.DataDir)
	if err != nil {
		return nil, err
	}
	store.db = db

	if err := store.setupRaft(); err != nil {
		return nil, err
	}

	return store, nil
}

// Join joins the node given ID, addr to the cluster
func (s *Store) Join(id, addr string) error {
	if s.getState() != raft.Leader {
		return raft.ErrNotLeader
	}

	configFuture := s.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		return err
	}

	for _, srv := range configFuture.Configuration().Servers {
		if srv.Address == raft.ServerAddress(addr) && srv.ID == raft.ServerID(id) {
			return nil
		}
		if srv.Address == raft.ServerAddress(addr) || srv.ID == raft.ServerID(id) {
			f := s.raft.RemoveServer(raft.ServerID(id), 0, 0)
			if f.Error() != nil {
				return f.Error()
			}
		}
	}

	f := s.raft.AddVoter(raft.ServerID(id), raft.ServerAddress(addr), 0, 0)

	return f.Error()
}

// GetOne gets next auto-increment ID for particular key
func (s *Store) GetOne(key string) (uint64, error) {
	if s.getState() != raft.Leader {
		return 0, raft.ErrNotLeader
	}

	cmd, err := newCommand(getOneCmd, &getOnePayload{
		Key: key,
	})
	if err != nil {
		return 0, err
	}

	cmdBytes, err := json.Marshal(cmd)
	if err != nil {
		return 0, err
	}

	f := s.raft.Apply(cmdBytes, applyTimeout)
	if f.Error() != nil {
		return 0, f.Error()
	}

	switch resp := f.Response().(type) {
	case *fsmGetOneResponse:
		return resp.value, resp.err
	case *fsmErrorResponse:
		return 0, resp.err
	default:
		return 0, errors.New("unknown error")
	}
}

// GetMany gets number of `quantity` of auto-increment ID for particular key
func (s *Store) GetMany(key string, quantity uint64) (uint64, uint64, error) {
	if s.getState() != raft.Leader {
		return 0, 0, raft.ErrNotLeader
	}

	cmd, err := newCommand(getManyCmd, &getManyPayload{
		Key:      key,
		Quantity: quantity,
	})
	if err != nil {
		return 0, 0, err
	}

	cmdBytes, err := json.Marshal(cmd)
	if err != nil {
		return 0, 0, err
	}

	f := s.raft.Apply(cmdBytes, applyTimeout)
	if f.Error() != nil {
		return 0, 0, f.Error()
	}

	switch resp := f.Response().(type) {
	case *fsmGetManyResponse:
		return resp.from, resp.to, resp.err
	case *fsmErrorResponse:
		return 0, 0, resp.err
	default:
		return 0, 0, errors.New("unknown error")
	}
}

// GetLastInserted gets the last inserted id for particular key. This API doesn't change database.
func (s *Store) GetLastInserted(key string) (uint64, error) {
	if s.getState() != raft.Leader {
		return 0, raft.ErrNotLeader
	}

	cmd, err := newCommand(getLastInsertedCmd, &getLastInsertedPayload{
		Key: key,
	})
	if err != nil {
		return 0, err
	}

	cmdBytes, err := json.Marshal(cmd)
	if err != nil {
		return 0, err
	}

	f := s.raft.Apply(cmdBytes, applyTimeout)
	if f.Error() != nil {
		return 0, f.Error()
	}

	switch resp := f.Response().(type) {
	case *fsmGetLastInsertedResponse:
		return resp.value, resp.err
	case *fsmErrorResponse:
		return 0, resp.err
	default:
		return 0, errors.New("unknown error")
	}
}

// Shutdown shutdowns the store
func (s *Store) Shutdown() error {
	s.db.Close()

	f := s.raft.Shutdown()
	if f.Error() != nil {
		return f.Error()
	}

	return nil
}

func (s *Store) setupRaft() error {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(s.config.NodeID)

	transport, err := raft.NewTCPTransport(s.config.RaftAddr, nil, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(s.config.RaftDir, 0777); err != nil {
		return err
	}

	store, err := raftboltdb.NewBoltStore(filepath.Join(s.config.RaftDir, "raft.db"))
	if err != nil {
		return err
	}

	logStore, err := raft.NewLogCache(logCacheCapacity, store)
	if err != nil {
		return err
	}

	stableStore := store

	snapshotStore, err := raft.NewFileSnapshotStore(s.config.RaftDir, retainSnapshotCount, os.Stderr)
	if err != nil {
		return err
	}

	ra, err := raft.NewRaft(config, newFSM(s.db), logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return err
	}
	s.raft = ra

	if s.config.Bootstrap {
		hasState, err := raft.HasExistingState(logStore, stableStore, snapshotStore)
		if err != nil {
			return err
		}
		if !hasState {
			configuration := raft.Configuration{
				Servers: []raft.Server{
					{
						ID:      config.LocalID,
						Address: transport.LocalAddr(),
					},
				},
			}
			ra.BootstrapCluster(configuration)
		}
	}

	if s.config.JoinAddr != "" {
		hasState, err := raft.HasExistingState(logStore, stableStore, snapshotStore)
		if err != nil {
			return err
		}
		if !hasState {
			if err := s.joinCluster(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Store) getState() raft.RaftState {
	return s.raft.State()
}

func (s *Store) joinCluster() error {
	conn, err := grpc.Dial(s.config.JoinAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewAutoIncrementClient(conn)

	if _, err := client.Join(context.Background(), &pb.JoinRequest{
		NodeID:      s.config.NodeID,
		NodeAddress: s.config.RaftAddr,
	}); err != nil {
		return err
	}

	return nil
}
