package main

import (
	"fmt"
	"time"
)

func main() {
	cc := make(chan chan int)
	for i := 1; i < 10+1; i++ {
		f := make(chan bool)
		go f1(cc, f)
		ch := <-cc
		ch <- i
		for sum := range ch {
			fmt.Print("Sum(", i, ")=", sum)
		}
		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum = sum + i
		}
		c <- sum
	case <-f:
		return
	}
}
