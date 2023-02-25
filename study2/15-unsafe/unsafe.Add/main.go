package main

import (
	"fmt"
	"unsafe"
)

// 此函数在一个(非安全)指针表示的地址上添加一个偏移量，然后返回一个新地址的指针‘
func main() {
	a := [16]int{3: 3, 9: 9, 11: 11}
	fmt.Println(a)
	eleSize := int(unsafe.Sizeof(a[0]))
	p9 := &a[9]
	up9 := unsafe.Pointer(p9)
	p3 := (*int)(unsafe.Add(up9, -6*eleSize))
	fmt.Println(*p3) // 3
	s := unsafe.Slice(p9, 5)[:3]
	fmt.Println(s)              // [9 0 11]
	fmt.Println(len(s), cap(s)) // 3 5

	t := unsafe.Slice((*int)(nil), 0)
	fmt.Println(t == nil) // true

	// 下面是两个不正确的调用。因为它们
	// 的返回结果引用了未知的内存块。
	_ = unsafe.Add(up9, 7*eleSize)
	_ = unsafe.Slice(p9, 8)
}
