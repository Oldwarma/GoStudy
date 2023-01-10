package main

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/zeromicro/go-zero/core/logx"
	"gostudy/study2/1-fileTran/conndata"
	"gostudy/tools/utils"
	"net"
	"runtime"
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
	req := new(conndata.Msg)
	_ = gconv.Struct(buf[:n], req)
	fmt.Println("req", req, string(buf[:n]))

	GetFile(req.Data.Info.Name, req.Data.Content)
}

func GetFile(fileName string, fileData []byte) error {
	var split []string
	var DirPath string
	if len(split) > 0 {
		fileName = split[0]
	}
	if runtime.GOOS == "windows" {
		DirPath = "D:/cep/"
	}
	//ioutil.WriteFile(DirPath+fileName, fileData, os.ModePerm)
	if !utils.WriteFile(DirPath+fileName, fileData) {
		return errors.New("文件写入失败")
	}

	savePathName := DirPath + fileName
	fmt.Println(savePathName)
	return nil
}
