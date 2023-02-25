package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//定义一个长度为16的int类型数组。并且给下标为3 9 11的元素赋值
	a := [16]int{3: 3, 9: 9, 11: 11}
	fmt.Println(a)
	//获取数组中每一个元素的字节大小
	eleSize := int(unsafe.Sizeof(a[0]))
	//获取下标为9的元素指针
	p9 := &a[9]
	//指针转换为unsafe.pointer类型的指针
	up9 := unsafe.Pointer(p9)
	fmt.Println("up9:", up9)
	//获取下标为3的元素值
	//unsafe.Add(up,-6*elseSize) 将下标为9的元素指针，向前移动六
	//(*int)转换为int类型的指针
	p3 := (*int)(unsafe.Add(up9, -6*eleSize))
	fmt.Println("*p3:", *p3)

	//截前五位，再截前三
	s := unsafe.Slice(p9, 5)[:3]
	fmt.Println("s:", s)
	fmt.Println("len:", len(s), "cap:", cap(s))

	t := unsafe.Slice((*int)(nil), 0)
	fmt.Println(t == nil, t)
}
