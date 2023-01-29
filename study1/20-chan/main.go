package main

import (
	"fmt"
	"math/rand"
)

var Close = false
var Data = make(map[int]bool)

func main() {
	A := make(chan int)
	B := make(chan int)
	go first(10, 20, A)
	go second(B, A)
	third(B)
}
func random(min, max int) int {
	return rand.Intn(max-min) + min
}
func first(min, max int, out chan<- int) {
	for {
		if Close {
			close(out)
			return
		}
		out <- random(min, max)
	}
}
func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Println(x, "x")
		_, ok := Data[x]
		if ok {
			Close = true
		} else {
			Data[x] = true
			out <- x
		}
	}
	fmt.Println(in)
	close(out)
}
func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Println("The sum of the random numbers is %d\n", sum)
}
func f1(c chan int, x int) {
	fmt.Println(x)
	c <- x
}
func f2(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
}
