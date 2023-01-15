package main

import "fmt"

func main() {
	fmt.Println(retThree(10))
}
func retThree(x int) (int, int, int) {
	return 2 * x, x * x, -x
}
