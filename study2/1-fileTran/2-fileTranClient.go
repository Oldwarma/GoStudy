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
	// 读取服务器回发的 OK

	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("发送信息失败，err:", err)
		return
	}
}
