package main

import (
	"context"
	"github.com/aibianma/rpclearning/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &proto.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(reply.GetValue())
}
