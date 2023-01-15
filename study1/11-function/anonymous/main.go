package main

import (
	"fmt"
	"strconv"
)

func main() {

	y, _ := strconv.Atoi("123654")

	square := func(s int) int {
		return s * s
	}
	fmt.Println("The squre of", y, "is", square(y))

	double := func(s int) int {
		return s + s
	}
	fmt.Println("The double of", y, "is", double(y))

	fmt.Println(doubleSquare(y))
	d, s := doubleSquare(y)
	fmt.Println(d, s)
}

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}
