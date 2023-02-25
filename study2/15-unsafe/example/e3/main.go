package main

import (
	"fmt"
	"unsafe"
)

// 底层切片结果体
type slice struct {
	array unsafe.Pointer // 指向底层数组的指针
	len   int            // slice 的长度
	cap   int            // slice 的容量
}

func main() {
	s := make([]int, 9, 20)

	l := (*int)(unsafe.Add(unsafe.Pointer(&s), uintptr(8)))
	println(*l)

	demo9()
}
func demo9() {
	m := make(map[int]string)
	l := (**int)(unsafe.Pointer(&m))
	fmt.Println(**l)
	fmt.Println(len(m))
}
