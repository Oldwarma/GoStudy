package main

import (
	"fmt"
	"unsafe"
)

// 该函数返回一个字符串，底层字节从ptr开始，
// 长度为len。len参数必须是证书类型或者是无类型常量，常量len参数必须是非负的并且可以用int类型的值来表示。
// 如果在运行时len为负，或者ptr为nil且len不为零，那么会发生panic
func main() {
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	//传进的字节不可修改
	s := unsafe.String(&b[0], len(b)) // hello
	fmt.Println(s)
	b[1] = 'A'
	fmt.Println(s)
	println(unsafe.String(&b[0], 3))
}
