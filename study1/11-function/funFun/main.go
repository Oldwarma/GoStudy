package main

import "fmt"

func main() {
	fmt.Println("function1", funFun(function1, 123))
	fmt.Println("function1", funFun(function2, 123))
	fmt.Println("function1", funFun(func(i int) int { return i * i }, 123))
}

func function1(i int) int {
	return i + i
}
func function2(i int) int {
	return i * i
}
func funFun(f func(int) int, v int) int {
	return f(v)
}
