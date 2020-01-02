package auto_increment

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/soheilhy/cmux"

	"github.com/ldmtam/raft-auto-increment/store"

	"github.com/ldmtam/raft-auto-increment/config"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"

	pb "github.com/ldmtam/raft-auto-increment/auto_increment/pb"
)

// Store represents the store interface
type Store interface {
	// Join joins the node with given ID and address
	Join(id, addr string) error

	// GetSingle gets next auto-increment ID for particular key
	GetSingle(key string) (uint64, error)

	// GetMultiple gets number of `quantity` of auto-increment ID for particular key
	GetMultiple(key string, quantity uint64) ([]uint64, error)

	// GetLast gets the last inserted id for particular key. This API doesn't change database.
	GetLast(key string) (uint64, error)

	// Shutdown shutdowns the store
	Shutdown() error
}

// AutoIncrement represents the AutoIncrement structure
type AutoIncrement struct {
	grpcServer *grpc.Server
	httpServer *http.Server

	store Store

	config *config.Config
}

// New returns new instance of AutoIncrement interface
func New(config *config.Config) (*AutoIncrement, error) {
	ai := &AutoIncrement{
		config: config,
	}

	store, err := store.New(config)
	if err != nil {
		return nil, err
	}
	ai.store = store

	listener, err := net.Listen("tcp", config.Addr)
	if err != nil {
		return nil, err
	}

	m := cmux.New(listener)
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	ai.grpcServer = grpc.NewServer()
	pb.RegisterAutoIncrementServer(ai.grpcServer, ai)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterAutoIncrementHandlerFromEndpoint(context.Background(), mux, ai.config.Addr, opts); err != nil {
		return nil, err
	}

	ai.httpServer = &http.Server{
		Handler: mux,
	}

	go ai.grpcServer.Serve(grpcListener)
	go ai.httpServer.Serve(httpListener)
	go m.Serve()

	return ai, nil
}

// Stop ...
func (ai *AutoIncrement) Stop() {
	ai.grpcServer.GracefulStop()
	ai.httpGracefulShutdown()
	ai.store.Shutdown()
}

func (ai *AutoIncrement) httpGracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ai.httpServer.Shutdown(ctx)
}
