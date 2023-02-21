package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"gostudy/study2/1-fileTran/conndata"
	"net"
	"os"
)

type Msg struct {
	Head Head
	Data Data
}

type Head struct {
	FunName   string
	Cot       string
	Mid       string
	Timestamp string
}
type Data struct {
	Info    interface{}
	Content []byte
}

func main() {
	// 主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	// 读取服务器回发的 OK
	req := &conndata.Msg{
		Head: conndata.Head{
			FunName: "DeviceUpgradeFileReq",
			Cot:     "1",
			Mid:     "2",
		},
		Data: conndata.Data{
			Info: conndata.Info{
				Name:     "12313.tar.gz",
				FileName: "0213ces.tar.gz",
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
	fmt.Println("发送完毕")
	for {

	}
}

func sendFile(conn net.Conn, filePath string) error {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return errors.New("文件不存在")
	}

	f, err := os.Open(filePath)
	if err != nil {
		return errors.New("打开文件失败")
	}
	defer f.Close()
	buf := make([]byte, fileInfo.Size())
	n, err := f.Read(buf)
	resp := FileTranLocal{}

	resp.FileInfo = fileInfo
	resp.Content = buf[:n]
	bs, _ := json.Marshal(resp)
	_, err = conn.Write(bs)
	if err != nil {
		return errors.New("文件写入失败")
	}
	return nil

}

type FileTranLocal struct {
	FileInfo os.FileInfo
	Content  []byte
}
