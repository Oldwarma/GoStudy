package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	bs, _ := json.Marshal("13213")
	openFile(bs)
}

func openFile(content []byte) bool {
	//没有就创建文件，后面是权限，返回文件的操作结构体*os.File,带权限
	file, err := os.OpenFile("D:/cep/openFile", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
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
			return true
		}
	}
	return false
}
