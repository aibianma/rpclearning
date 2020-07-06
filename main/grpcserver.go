package main

import (
	"github.com/aibianma/rpclearning/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, new(proto.HelloServiceImpl))
	lis, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
