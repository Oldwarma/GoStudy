package main

import "fmt"

func main() {
	fmt.Println(degui(3, 5))
}
func degui(a, b int) int {
	c := a + b
	if c < 100 {
		fmt.Println(c, a, b)
		return degui(c, a)
	}
	return c
}
