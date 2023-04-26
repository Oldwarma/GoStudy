package main

import (
	"fmt"
	"net"
)

func main() {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:5178")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	for {
		pro()
	}
}

func pro() {

	fmt.Println("123")
	for {

	}
}
