package main

import (
	"fmt"
	"os"
)

func writeFile(s string) {
	file, err := os.OpenFile("D://text.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("打开文件错误", err)
	}
	defer file.Close()
	if err != nil {
		fmt.Println(file, err)
		return
	}
	file.WriteString(s)
	file.Sync()
	fmt.Println("文件存入")
}
func main() {
	writeFile("你干嘛哎呦")
}
