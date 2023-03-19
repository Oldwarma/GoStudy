package main

import (
	"fmt"
	"google.golang.org/grpc"
	"gostudy/study3/7-Grpc/server/rpc"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatalf("启动grpc server失败")
		return
	}
	grpcServer := grpc.NewServer()

	rpc.RegisterServerServer(grpcServer, Server{})

	log.Println("service start")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("启动grpc server失败")
	}
}
