package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ma = make(map[int]chan string)
	mu sync.Mutex
	ch = make(chan string, 1)
)

func main() {
	ma[1] = ch
	go chan1()
	func() {
		for {
			select {
			case msg := <-ma[1]:
				fmt.Println(msg)
			case <-time.After(time.Second * 3):
				fmt.Println("超时")
			}
		}
	}()
	ma[1] <- "你干嘛哎呦2"

}
func chan1() {

	ma[1] <- "你干嘛哎呦1"

}
