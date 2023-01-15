package main

import "fmt"

func main() {
	text()
}

//copy,前面参数被后面替换
func text() {
	a6 := []int{10, 11, 12, 13, 14}
	a4 := []int{-1, -2, -3, -4}

	n := copy(a6, a4)
	fmt.Printf("n=%v,a6=%v,a4=%v \n", n, a6, a4)
	//n := copy(a4, a6)
	//fmt.Printf("n=%v,a6=%v,a4=%v", n, a6, a4)
}
