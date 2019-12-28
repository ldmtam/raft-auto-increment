package auto_increment

import (
	"context"
	"net"
	"net/http"
	"time"

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
