package main

import (
	"fmt"
	"time"
)

func worker1(id int, in chan bool, out chan bool) {

}

func workers2(id int, in <-chan bool, out chan<- bool) {
	fmt.Printf("worker %d start\n", id)
	<-in //阻塞
	fmt.Printf("worker %d quit\n", id)
	out <- true
}
func workers3(id int, in <-chan bool, out chan<- bool) {
	fmt.Printf("worker %d start\n", id)
	for {
		fmt.Printf("worker %d doing\n", id)
		time.Sleep(200 * time.Microsecond)
		select {
		case <-in:
			out <- true
			return
		default:
		}
	}
}

/*
chan<-只写
<-chan 只读
*/
func main() {
	in := make(chan bool)
	out := make(chan bool)
	workers := 3

	for i := 0; i < workers; i++ {
		go workers3(i, in, out)
	}
	go func() {
		time.Sleep(time.Second)
		close(in)
		//for i := 0; i < workers; i++ {
		//	in <- true
		//}
	}()
	count := 0
	for count < workers {
		<-out //阻塞
		count++
	}
	fmt.Println("ok")

}
