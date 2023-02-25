package main

import "unsafe"

// 返回类型x所占的字节数，但不包含x所指向的内容大小，对于一个slice来说返回的就是slice header的大小
func main() {
	a := int(1)
	b := float32(2.0)
	s := "Ethan"
	sl := make([]string, 16)

	println(unsafe.Sizeof(a))  // 8
	println(unsafe.Sizeof(b))  // 4
	println(unsafe.Sizeof(s))  // 16
	println(unsafe.Sizeof(sl)) // 24
}
