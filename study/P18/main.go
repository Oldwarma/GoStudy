package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type huaji struct {
	name string
	age  int
}

func GenerateintA(done chan huaji) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan huaji)
	ch := GenerateintA(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(runtime.NumGoroutine())
}
