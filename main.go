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
)

func init() {
	flag.StringVar(&httpAddr, "http-addr", "localhost:3000", "Service http address")
	flag.StringVar(&grpcAddr, "grpc-addr", "localhost:4000", "Service grpc address")
}

func main() {
	flag.Parse()

	config := &autoIncrement.Config{
		HttpAddr: httpAddr,
		GrpcAddr: grpcAddr,
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
