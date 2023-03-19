package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gostudy/study3/7-Grpc/client/clientRpc"
	"log"
)

// 在使用 grpc.WithInsecure() 创建连接时，所有的通信都是明文传输的
func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	ServerClient := clientRpc.NewServerClient(conn)

	hello, err := ServerClient.Hello(context.Background(), &clientRpc.Empty{})
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	log.Println(hello, err)
}
