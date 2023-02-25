package main

import (
	"fmt"
	st "gostudy/example/e1/cst"
)

func main() {
	ch := make(chan int)
	st.ReadMap[1] = ch
	pro()
	for {

	}
}

func pro() {
	ch := st.ReadMap[1]
	select {
	case msg := <-ch:
		fmt.Println("msg==", msg)

	}

}
