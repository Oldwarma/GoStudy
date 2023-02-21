package main

import (
	"fmt"
	st "gostudy/text/text1/cst"
	"time"
)

func main() {
	ch := make(chan int, 10)
	st.ReadMap[1] = ch

	pro1()

	pro2()

	for {

	}
}
func pro1() {
	ch := st.ReadMap[1]
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("iii", i)
	}
	fmt.Println("发送完毕")
	close(ch)
}
func pro2() {
	ch := st.ReadMap[1]
	select {
	case msg := <-ch:
		fmt.Println("msg==", msg)
		pro2()
	case <-time.After(time.Second * 3):
		fmt.Println("超时")
	}

}
