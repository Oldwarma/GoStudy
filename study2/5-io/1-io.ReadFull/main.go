package main

import (
	"fmt"
	"gostudy/tools/byteoperation"
	"io"
	"strings"
)

//当进行readfull读取reader数据时，reader中有一个隐形指针指向最后读到数据的后一位，
//当下一次进行readfull时从当前指针的位置进行读取，即前一次读取的数据不再进行读取
func main() {
	reader := strings.NewReader("Geeks")

	buffer := make([]byte, 4)

	n, err := io.ReadFull(reader, buffer)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of bytes in the buffer: %d\n", n)
	fmt.Printf("Content in buffer: %s\n", buffer)

	buffer2 := make([]byte, 1) //超过会报错
	n, err = io.ReadFull(reader, buffer2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of bytes in the buffer: %d\n", n)
	fmt.Printf("Content in buffer: %s\n", buffer2)

	fmt.Printf(string(byteoperation.RandBytes(2)))
}
