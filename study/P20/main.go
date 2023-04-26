package main

import (
	"fmt"
	"time"
)

//
//import (
//	"fmt"
//	"math/rand"
//)
//
//func e23()  {
//	ch := GenerateInt()
//	for i := 0;i<100;i++ {
//		fmt.Println(<- ch)
//	}
//}
//
//
//func GenerateInt() chan int {
//	ch := make(chan int,20)
//	go func() {
//		for  {
//			select {
//			case ch <- <- GenerateintA():
//				fmt.Println("拉拉1")
//			case ch <- <- GenerateintB():
//
//				fmt.Println("拉拉2")
//			}
//		}
//	}()
//	return ch
//}
//
//
//func GenerateintA() chan int {
//	ch := make(chan int,10)
//	go func() {
//		for {
//			ch <- rand.Int()
//		}
//	}()
//
//	return ch
//}
//
//
//func GenerateintB() chan int {
//	ch := make(chan int,10)
//	go func() {
//		for {
//			ch <- rand.Int()
//		}
//	}()
//
//	return ch
//}

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		fmt.Println("阻塞")
		for v := range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		for i := 0; i < 10; i++ {
			in <- i
		}
		//close(in)
		//time.Sleep(time.Second * 3)
		in <- 50
		//close(in)
	}()

	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Println(v)
	}

}
