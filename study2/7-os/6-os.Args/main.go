package main

import (
	"fmt"
	"os"
)

//os.Stdin 命令行输入
//命令行参数
func main() {
	args()
}

func args() {
	a := os.Args
	for _, v := range a {
		fmt.Println(v)
	}
}
