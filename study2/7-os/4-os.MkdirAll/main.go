package main

import (
	"fmt"
	"os"
)

func main() {
	mkdirAll()
}

//创建目录，必须存在
func mkdirAll() {
	//os.Mkdir()创建单个目录
	err := os.MkdirAll("D:/cep/text0112/1321", os.ModePerm)
	if err != nil {
		fmt.Println("create error", err)
	}
	fmt.Println("ok")
}
