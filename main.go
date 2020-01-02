package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	autoIncrement "github.com/ldmtam/raft-auto-increment/auto_increment"
	"github.com/ldmtam/raft-auto-increment/config"
)

var (
	addr    string
	dataDir string

	raftAddr  string
	joinAddr  string
	raftDir   string
	bootstrap bool
	nodeID    string
)

func init() {
	flag.StringVar(&addr, "addr", "localhost:3000", "Service grpc/http address")
	flag.StringVar(&dataDir, "data-dir", "./data", "Data directory")

	flag.StringVar(&raftAddr, "raft-addr", "localhost:5000", "Raft address")
	flag.StringVar(&joinAddr, "join-addr", "localhost:3000", "Leader address")
	flag.StringVar(&raftDir, "raft-dir", "./raft", "Raft directory")
	flag.BoolVar(&bootstrap, "bootstrap", false, "Start as bootstrap node")
	flag.StringVar(&nodeID, "id", "1", "Raft node ID")
}

func main() {
	flag.Parse()

	config := &config.Config{
		NodeID:    nodeID,
		RaftAddr:  raftAddr,
		RaftDir:   raftDir,
		Bootstrap: bootstrap,
		Addr:      addr,
		DataDir:   dataDir,
	}

	service, err := autoIncrement.New(config)
	if err != nil {
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
