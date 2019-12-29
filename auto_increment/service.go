package auto_increment

import (
	"context"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	raftboltdb "github.com/hashicorp/raft-boltdb"

	"github.com/hashicorp/raft"
	"github.com/ldmtam/raft-auto-increment/auto_increment/database/leveldb"

	"github.com/ldmtam/raft-auto-increment/auto_increment/database"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"

	pb "github.com/ldmtam/raft-auto-increment/auto_increment/pb"
)

// AutoIncrement represents the AutoIncrement service interface
type AutoIncrement interface {
	Start() error
	Stop()
}

type autoIncrement struct {
	grpcServer *grpc.Server
	httpServer *http.Server

	db database.AutoIncrement

	raft *raft.Raft

	config *Config
}

// New returns new instance of AutoIncrement interface
func New(config *Config) AutoIncrement {
	return &autoIncrement{
		config: config,
	}
}

func (s *autoIncrement) Start() error {
	grpcListener, err := net.Listen("tcp", s.config.GrpcAddr)
	if err != nil {
		return err
	}

	httpListener, err := net.Listen("tcp", s.config.HttpAddr)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()
	pb.RegisterAutoIncrementServer(s.grpcServer, s)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterAutoIncrementHandlerFromEndpoint(context.Background(), mux, s.config.GrpcAddr, opts); err != nil {
		return err
	}

	s.httpServer = &http.Server{
		Handler: mux,
	}

	s.db, err = leveldb.New(s.config.DataDir)
	if err != nil {
		return err
	}

	if err := s.setupRaft(); err != nil {
		return err
	}

	go s.grpcServer.Serve(grpcListener)
	go s.httpServer.Serve(httpListener)

	return nil
}

func (s *autoIncrement) Stop() {
	s.grpcServer.GracefulStop()
	s.httpGracefulShutdown()
	s.db.Close()
}

func (s *autoIncrement) httpGracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.httpServer.Shutdown(ctx)
}

func (s *autoIncrement) setupRaft() error {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(s.config.NodeID)

	transport, err := raft.NewTCPTransport(s.config.RaftAddr, nil, 3, 10*time.Second, os.Stderr)
	if err != nil {
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
