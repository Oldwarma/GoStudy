package main

import "unsafe"

type User struct {
	Name     string
	Nickname string
	Age      int
}

// 返回结构体成员在内存中位置离结构体起始处的字节数，所传递的参数必须是结构体的成员
func main() {
	u := User{
		Age:      10,
		Name:     "Ethan",
		Nickname: "Leo",
	}
	println(unsafe.Offsetof(u.Age))
	println(unsafe.Offsetof(u.Nickname))
	println(unsafe.Offsetof(u.Name))
}
