package main

import "fmt"

var (
	ma = make(map[int]chan int)
	ch = make(chan int)
)

//读与写占用主线程时，另一个应该在主线程之前开启协程，或者都开协程
func main() {
	ma[1] = ch
	go send()
	read()

}

func read() {
	for {
		select {
		case msg := <-ma[1]:
			fmt.Println(msg)
		}
	}

}

func send() {

	for i := 0; i < 10; i++ {
		if i > 5 {
			ma[1] <- i
		}
	}
	for {
	}

}
