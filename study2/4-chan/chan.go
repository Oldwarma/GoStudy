package main

import (
	"fmt"
	"time"
)

var (
	ch = make(chan string, 2)
)

func main() {
	go chan1()
	go chan2()
	go chan3()
	for true {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-time.After(time.Second * 3):
			fmt.Println("超时")
		}
	}

}

func chan1() {
	ch <- "你干嘛哎呦"
}
func chan2() {
	ch <- "只因你实在是太美"
}
func chan3() {
	ch <- "哦，baby"
}
