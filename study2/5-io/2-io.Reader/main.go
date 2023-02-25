package main

import (
	"fmt"
	"log"
	"strings"
)

// io 提供 I/O 原语的基本接口。在 Go 语言标准库 strings、bytes、bufio、和 os 中，都有实现 io.Reader 的类型
func main() {
	myReader := strings.NewReader("Hello, world!")
	buffer := make([]byte, 5)
	n, err := myReader.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s", n, buffer[:n])

}
