package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(TwoReverse())
	}
}

// 方法一为
var seq int = 0

func OneReverse() func() int {
	return func() int {
		res := seq
		seq = (seq + 1) % 2
		return res
	}
}

// 方法二为
func TwoReverse() int {
	seqGen := OneReverse()
	return seqGen()
}
