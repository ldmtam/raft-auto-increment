package auto_increment

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/soheilhy/cmux"

	"github.com/ldmtam/raft-auto-increment/store"

	"github.com/ldmtam/raft-auto-increment/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ldmtam/raft-auto-increment/pb"
)

// Store represents the store interface
type Store interface {
	// Join joins the node with given ID and address
	Join(id, addr string) error

	// GetOne gets next auto-increment ID for particular key
	GetOne(key string) (uint64, error)

	// GetMany gets number of `quantity` of auto-increment ID for particular key
	GetMany(key string, quantity uint64) (uint64, uint64, error)

	// GetLastInserted gets the last inserted id for particular key. This API doesn't change database.
	GetLastInserted(key string) (uint64, error)

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

	listener, err := net.Listen("tcp", config.NodeAddr)
	if err != nil {
		return nil, err
	}

	m := cmux.New(listener)
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	ai.grpcServer = grpc.NewServer()
	pb.RegisterAutoIncrementServiceServer(ai.grpcServer, ai)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	if err := pb.RegisterAutoIncrementServiceHandlerFromEndpoint(context.Background(), mux, ai.config.NodeAddr, opts); err != nil {
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
	ai.httpGracefulShutdown()
	ai.store.Shutdown()
}

func (ai *AutoIncrement) httpGracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ai.httpServer.Shutdown(ctx)
}
