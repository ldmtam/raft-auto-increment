package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	gobitcask "github.com/ldmtam/go-bitcask"
	"github.com/ldmtam/raft-auto-increment/common"
	raftbitcask "github.com/ldmtam/raft-bitcask"

	"github.com/ldmtam/raft-auto-increment/database/memdb"

	pb "github.com/ldmtam/raft-auto-increment/pb"

	raftbadgerdb "github.com/BBVA/raft-badger"
	"github.com/hashicorp/raft"
	"github.com/ldmtam/raft-auto-increment/config"
	"github.com/ldmtam/raft-auto-increment/database"
	"google.golang.org/grpc"
)

const (
	retainSnapshotCount = 2
	logCacheCapacity    = 512
	applyTimeout        = 10 * time.Second
)

type leaderInfo struct {
	NodeAddr string
	RaftAddr string
	RaftID   string
}

// Store represents the Store structure
type Store struct {
	raft *raft.Raft

	leader     *leaderInfo
	leaderConn *grpc.ClientConn

	shutdownCh chan struct{}
	wg         sync.WaitGroup

	db     database.AutoIncrement
	config *config.Config
}

// New return new instance of Store object
func New(config *config.Config) (*Store, error) {
	var err error

	store := &Store{
		config:     config,
		leader:     new(leaderInfo),
		shutdownCh: make(chan struct{}),
	}

	store.db, err = memdb.New(nil)
	if err != nil {
		return nil, err
	}

	if err := store.setupRaft(); err != nil {
		return nil, err
	}

	store.wg.Add(1)
	go store.monitorLeadership()

	return store, nil
}

// Join joins the node given ID, addr to the cluster
func (s *Store) Join(id, addr string) error {
	if f := s.raft.VerifyLeader(); f.Error() != nil {
		return s.forwardJoinRequest(id, addr)
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
	if f := s.raft.VerifyLeader(); f.Error() != nil {
		return s.forwardGetOneRequest(key)
	}

	cmd, err := newCommand(setIDCmd, &setIDPayload{
		Key:      key,
		Quantity: 1,
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

	raftResp, ok := f.Response().(*fsmResponse)
	if !ok {
		return 0, errors.New("unknown error")
	}
	if raftResp.err != nil {
		return 0, err
	}

	getOneResp, ok := raftResp.resp.(*getOneIDResponse)
	if !ok {
		return 0, errors.New("unknown error")
	}

	return getOneResp.value, nil
}

// GetMany gets number of `quantity` of auto-increment ID for particular key
func (s *Store) GetMany(key string, quantity uint64) (uint64, uint64, error) {
	if f := s.raft.VerifyLeader(); f.Error() != nil {
		return s.forwardGetManyRequest(key, quantity)
	}

	cmd, err := newCommand(setIDCmd, &setIDPayload{
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

	raftResp, ok := f.Response().(*fsmResponse)
	if !ok {
		return 0, 0, errors.New("unknown error")
	}
	if raftResp.err != nil {
		return 0, 0, raftResp.err
	}

	getManyResp, ok := raftResp.resp.(*getManyIDsResponse)
	if !ok {
		return 0, 0, errors.New("unknown error")
	}

	return getManyResp.from, getManyResp.to, nil
}

// GetLastInserted gets the last inserted id for particular key. This API doesn't change database.
func (s *Store) GetLastInserted(key string) (uint64, error) {
	if f := s.raft.VerifyLeader(); f.Error() != nil {
		return s.forwardGetLastInsertedRequest(key)
	}

	return s.db.GetLastInserted(key)
}

func (s *Store) setLeaderInfo() error {
	if f := s.raft.VerifyLeader(); f.Error() != nil {
		return raft.ErrNotLeader
	}

	cmd, err := newCommand(setLeaderInfoCmd, &setLeaderInfoPayload{
		NodeAddr: s.config.NodeAddr,
		RaftAddr: s.config.RaftAddr,
		RaftID:   s.config.RaftID,
	})
	if err != nil {
		return err
	}

	cmdBytes, err := json.Marshal(cmd)
	if err != nil {
		return err
	}

	f := s.raft.Apply(cmdBytes, applyTimeout)
	if f.Error() != nil {
		return f.Error()
	}

	resp, ok := f.Response().(*fsmResponse)
	if !ok {
		return errors.New("unknown error")
	}

	return resp.err
}

// Shutdown shutdowns the store
func (s *Store) Shutdown() error {
	if s.leaderConn != nil {
		s.leaderConn.Close()
	}

	s.db.Close()

	close(s.shutdownCh)
	s.wg.Wait()

	f := s.raft.Shutdown()
	if f.Error() != nil {
		return f.Error()
	}

	return nil
}

func (s *Store) setupRaft() error {
	config := raft.DefaultConfig()
	config.LogLevel = "INFO"
	config.LocalID = raft.ServerID(s.config.RaftID)

	transport, err := raft.NewTCPTransport(s.config.RaftAddr, nil, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(s.config.RaftDir, 0777); err != nil {
		return err
	}

	var stableStore raft.StableStore
	var logStore raft.LogStore

	switch s.config.Storage {
	case common.BADGER_STORAGE:
		store, err := raftbadgerdb.NewBadgerStore(s.config.RaftDir)
		if err != nil {
			return err
		}

		logStore, err = raft.NewLogCache(logCacheCapacity, store)
		if err != nil {
			return err
		}

		stableStore = store
	case common.BITCASK_STORAGE:
		store, err := raftbitcask.New(
			gobitcask.WithDirName(s.config.RaftDir),
			gobitcask.WithSegmentSize(256*1024*1024), // 128MB segment size
			gobitcask.WithMergeOpt(&gobitcask.MergeOption{
				MinFiles: 5,
				Interval: 6 * time.Hour,
			}),
		)
		if err != nil {
			return err
		}

		logStore, err = raft.NewLogCache(logCacheCapacity, store)
		if err != nil {
			return err
		}

		stableStore = store
	default:
		return common.ErrStorageNotAvailable
	}

	snapshotStore, err := raft.NewFileSnapshotStore(s.config.RaftDir, retainSnapshotCount, os.Stderr)
	if err != nil {
		return err
	}

	ra, err := raft.NewRaft(config, newFSM(s.db, s), logStore, stableStore, snapshotStore, transport)
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

func (s *Store) monitorLeadership() {
	defer s.wg.Done()
	for {
		select {
		case leader := <-s.raft.LeaderCh():
			if leader {
				if err := s.setLeaderInfo(); err != nil {
					fmt.Println(err)
				}
			}
		case <-s.shutdownCh:
			return
		}
	}
}

func (s *Store) joinCluster() error {
	conn, err := grpc.Dial(s.config.JoinAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewAutoIncrementServiceClient(conn)

	if _, err := client.Join(context.Background(), &pb.JoinRequest{
		RaftID:      s.config.RaftID,
		RaftAddress: s.config.RaftAddr,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Store) forwardJoinRequest(raftID, raftAddr string) error {
	var err error

	if s.leaderConn == nil {
		s.leaderConn, err = grpc.Dial(s.leader.NodeAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
	}

	client := pb.NewAutoIncrementServiceClient(s.leaderConn)

	if _, err := client.Join(context.Background(), &pb.JoinRequest{
		RaftID:      raftID,
		RaftAddress: raftAddr,
	}); err != nil {
		return err
	}

	return nil
}

func (s *Store) forwardGetOneRequest(key string) (uint64, error) {
	var err error

	if s.leaderConn == nil {
		s.leaderConn, err = grpc.Dial(s.leader.NodeAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return 0, err
		}
	}

	client := pb.NewAutoIncrementServiceClient(s.leaderConn)

	resp, err := client.GetOne(context.Background(), &pb.GetOneRequest{Key: key})
	if err != nil {
		return 0, err
	}

	return resp.Value, nil
}

func (s *Store) forwardGetManyRequest(key string, quantity uint64) (uint64, uint64, error) {
	var err error

	if s.leaderConn == nil {
		s.leaderConn, err = grpc.Dial(s.leader.NodeAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return 0, 0, err
		}
	}

	client := pb.NewAutoIncrementServiceClient(s.leaderConn)

	resp, err := client.GetMany(context.Background(), &pb.GetManyRequest{
		Key:      key,
		Quantity: quantity,
	})
	if err != nil {
		return 0, 0, err
	}

	return resp.From, resp.To, nil
}

func (s *Store) forwardGetLastInsertedRequest(key string) (uint64, error) {
	var err error

	if s.leaderConn == nil {
		s.leaderConn, err = grpc.Dial(s.leader.NodeAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return 0, err
		}
	}

	client := pb.NewAutoIncrementServiceClient(s.leaderConn)

	resp, err := client.GetLastInserted(context.Background(), &pb.GetLastInsertedRequest{Key: key})
	if err != nil {
		return 0, err
	}

	return resp.Value, nil
}
