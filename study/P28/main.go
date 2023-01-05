package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	allList = make(chan string, 10)
)

func main() {
	go func() {
		HandleChannel()
	}()
	for i := 0; i < 100; i++ {
		c := strconv.FormatInt(int64(i), 10)
		allList <- c
		fmt.Println("传入:", c)
	}
}

func HandleChannel() {
	for {
		select {
		case data := <-allList:
			fmt.Println(data)
			time.Sleep(time.Second)
		case <-time.After(time.Second * 2):
			fmt.Println("超时")
		}
	}
}
