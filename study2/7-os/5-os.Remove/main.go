package main

import (
	"fmt"
	"os"
)

func main() {
	remove()
}

//删除文件,需要存在
func remove() {
	err := os.Remove("D:/cep/text0112")
	//err := os.RemoveAll("D:/cep/text0112/0112")
	if err != nil {
		fmt.Println("create error", err)
	}
	fmt.Println("ok")
}
