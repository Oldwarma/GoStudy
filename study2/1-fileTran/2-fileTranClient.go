package main

import (
	"encoding/json"
	"fmt"
	"gostudy/study2/1-fileTran/conndata"
	"net"
	"os"
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
	req := &conndata.Msg{
		Head: conndata.Head{
			FunName: "asdas",
			Cot:     "1",
			Mid:     "2",
		},
		Data: conndata.Data{
			Info: conndata.Info{
				Name: "12313.tar.gz",
			},
		},
	}
	fileInfo, err := os.Stat("D:/cep/0110.tar.gz")
	if err != nil {
		fmt.Println("文件不存在")
	}
	f, err := os.Open("D:/cep/0110.tar.gz")
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, fileInfo.Size())
	n, err := f.Read(buf)
	req.Data.Content = buf[:n]
	bs, _ := json.Marshal(req)
	_, err = conn.Write(bs)
	if err != nil {
		fmt.Println("发送信息失败，err:", err)
		return
	}
}
