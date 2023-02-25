package main

import (
	"fmt"
	"unsafe"
)

// StringData返回一个指向str的底层字节的指针。
// 对于空字符串，返回值未指定，可能为nil。由于Go字符串是不可变的，因此不能修改StringData返回的字节。
func main() {
	s := "hello"
	t := unsafe.Slice(unsafe.StringData(s), len(s))
	fmt.Println(t)
}
