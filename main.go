package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	autoIncrement "github.com/ldmtam/raft-auto-increment/auto_increment"
)

var (
	httpAddr string
	grpcAddr string
	dataDir  string

	raftAddr  string
	raftDir   string
	bootstrap bool
	nodeID    string
)

func init() {
	flag.StringVar(&httpAddr, "http-addr", "localhost:3000", "Service http address")
	flag.StringVar(&grpcAddr, "grpc-addr", "localhost:4000", "Service grpc address")
	flag.StringVar(&dataDir, "data-dir", "./data", "Data directory")

	flag.StringVar(&raftAddr, "raft-addr", "localhost:5000", "Raft address")
	flag.StringVar(&raftDir, "raft-dir", "./raft", "Raft directory")
	flag.BoolVar(&bootstrap, "bootstrap", false, "Start as bootstrap node")
	flag.StringVar(&nodeID, "id", "1", "Raft node ID")
}

func main() {
	flag.Parse()

	config := &autoIncrement.Config{
		NodeID:    nodeID,
		RaftAddr:  raftAddr,
		RaftDir:   raftDir,
		Bootstrap: bootstrap,
		HttpAddr:  httpAddr,
		GrpcAddr:  grpcAddr,
		DataDir:   dataDir,
	}

	service := autoIncrement.New(config)

	if err := service.Start(); err != nil {
		panic(err)
	}

	waitExit()

	service.Stop()
}

func waitExit() {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-termChan
}
