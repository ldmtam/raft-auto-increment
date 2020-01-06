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
	nodeAddr string
	dataDir  string

	raftAddr  string
	joinAddr  string
	raftDir   string
	bootstrap bool
	raftID    string
)

func init() {
	flag.StringVar(&nodeAddr, "node-addr", "localhost:3000", "Service grpc/http address")
	flag.StringVar(&dataDir, "data-dir", "./data", "Data directory")

	flag.StringVar(&raftAddr, "raft-addr", "localhost:5000", "Raft address")
	flag.StringVar(&joinAddr, "join-addr", "", "Leader address")
	flag.StringVar(&raftDir, "raft-dir", "./raft", "Raft directory")
	flag.BoolVar(&bootstrap, "bootstrap", false, "Start as bootstrap node")
	flag.StringVar(&raftID, "raft-id", "1", "Raft node ID")
}

func main() {
	flag.Parse()

	config := &config.Config{
		RaftID:    raftID,
		RaftAddr:  raftAddr,
		RaftDir:   raftDir,
		Bootstrap: bootstrap,
		JoinAddr:  joinAddr,
		NodeAddr:  nodeAddr,
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
