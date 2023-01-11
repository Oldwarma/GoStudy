package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	bs, _ := json.Marshal("1312312312312312213")
	open(bs)
}

func open(content []byte) bool {
	//需要已有文件，返回文件的操作结构体*os.File,不带权限，且没有权限
	file, err := os.Open("D:/cep/openFile")
	if err != nil {
		panic(err)
	}
	//os的open前缀要close
	defer func() {
		file.Close()
	}()
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {

		write, err := file.Write(content)
		if err == nil && write > 0 {
			fmt.Println("写入失败")
			return true
		}
	}
	return false
}
