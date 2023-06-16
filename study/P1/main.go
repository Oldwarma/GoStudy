package main

import (
	"runtime"
)

func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		c <- struct{}{}
	}(c, ci)
	println("NUM:", runtime.NumGoroutine())
	<-c
	println("NUM:", runtime.NumGoroutine())

	println("NUM:", runtime.NumGoroutine())

	for v := range ci {
		println(v)
	}
}
