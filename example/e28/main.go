package main

import (
	"fmt"
)

func main() {
	fmt.Println("wasm测试")
	for i := 10; i > 0; i-- {
		fmt.Printf("ajie%v天后暴富\n", i)
	}
	c := make(chan struct{}, 0)

	// 注册add函数
	//js.Global().Set("add", js.FuncOf(Add1))

	<-c
}

func Add1(a, b int) {
	fmt.Println(a + b)
}
