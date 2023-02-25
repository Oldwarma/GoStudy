package main

import "unsafe"

// 返回的值是指当类型进行内存对齐时，它分配到的内存地址可以整除该函数的返回值
func main() {
	a := 1
	b := int8(8)
	s := "1"
	println(unsafe.Alignof(a))
	println(unsafe.Alignof(b))
	println(unsafe.Alignof(s))
}

//以上三个函数返回的结果都是uintptr类型，这个类型可以和unsafe.Pointer类型相互转换。
//由于三个函数都在编译期间执行，所以它们的结果可以直接赋值给const变量
