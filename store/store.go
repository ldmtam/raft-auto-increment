package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"github.com/ldmtam/raft-auto-increment/config"
	"github.com/ldmtam/raft-auto-increment/database"
	"github.com/ldmtam/raft-auto-increment/database/leveldb"
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

	db, err := leveldb.New(store.config.DataDir)
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
	return nil
}

// GetSingle gets next auto-increment ID for particular key
func (s *Store) GetSingle(key string) (uint64, error) {
	cmd, err := newCommand(getSingleCmd, &getSinglePayLoad{
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

	resp := f.Response().(*getSingleResponse)

	return resp.value, resp.err
}

// GetMultiple gets number of `quantity` of auto-increment ID for particular key
func (s *Store) GetMultiple(key string, quantity uint64) ([]uint64, error) {
	cmd, err := newCommand(getMultipleCmd, &getMultiplePayload{
		Key:      key,
		Quantity: quantity,
	})
	if err != nil {
		return nil, err
	}

	cmdBytes, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}

	f := s.raft.Apply(cmdBytes, applyTimeout)
	if f.Error() != nil {
		return nil, f.Error()
	}

	resp := f.Response().(*getMultipleResponse)

	return resp.values, resp.err
}

// GetLast gets the last inserted id for particular key. This API doesn't change database.
func (s *Store) GetLast(key string) (uint64, error) {
	cmd, err := newCommand(getLastCmd, &getLastPayload{
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

	resp := f.Response().(*getLastResponse)

	return resp.value, resp.err
}

// Shutdown shutdowns the store
func (s *Store) Shutdown() error {
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

	ra, err := raft.NewRaft(config, nil, logStore, stableStore, snapshotStore, transport)
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

	return nil
}
