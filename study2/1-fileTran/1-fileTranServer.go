package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:5178")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer l.Close()
	fmt.Println("服务端启动成功，等待发送端发送文件")
	for {
		conn, accErr := l.Accept()
		if accErr != nil {
			logx.Errorf("对接文件tcp链接失败:%s", accErr.Error())
		}

		go process(conn)
	}

	return
}

func process(conn net.Conn) {
	defer conn.Close()
	//定义接收信息的字节数组
	var buf [10241024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("获取信息失败")
		return
	}
	fmt.Println("一起吃饭吗？", string(buf[:n]))
}
