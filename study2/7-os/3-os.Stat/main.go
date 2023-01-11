package main

import (
	"fmt"
	"os"
)

func main() {
	stat()
}
func stat() {
	file, err := os.Stat("D:/cep/openFile")
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}
