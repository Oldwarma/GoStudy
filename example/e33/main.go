package main

import "fmt"

func main() {
	arr := []byte{1, 2, 3}
	var goodsStatusOk int
	for range arr {
		goodsStatusOk++
	}
	fmt.Println(goodsStatusOk)
}
