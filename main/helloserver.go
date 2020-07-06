package main

import (
	"github.com/aibianma/rpclearning/hello"
	log "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(hello.HelloService))
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	rpc.ServeConn(conn)
}
