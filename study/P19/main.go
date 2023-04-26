package main

import (
	"fmt"
	"math/rand"
)

//import (
//	"fmt"
//	"time"
//)
//
//func e23() {
//	println("start e23")
//	cond1 := make(chan int)
//	cond2 := make(chan uint64)
//
//	go func() {
//		for i := 0; ; i++ {
//			cond1 <- i
//		}
//	}()
//
//	go func() {
//		var i uint64
//		for ; ; i++ {
//			cond2 <- i
//		}
//	}()
//
//	endCond := false
//	for endCond != true {
//		select {
//		case a := <-cond1:
//			if a > 99 {
//				println("end with cond1")
//				endCond = true
//			}else {
//				fmt.Println(a,"滑稽1")
//			}
//		case b := <-cond2:
//			if b == 100 {
//				println("end with cond2")
//				endCond = true
//			}else {
//				fmt.Println(b,"滑稽2")
//			}
//		case <-time.After(time.Microsecond):
//			println("end with clienttimeout")
//			endCond = true
//		}
//	}
//	println("end e23")
//}

//---------------------------------------------

type Student struct {
	Id   int64
	Name string
}

//var (
//	cMap = make(map[int64]chan Student)
//)
//
//func e23()  {
//	defer func() {
//		e := recover()
//		if e != nil  {
//
//		}
//	}()
//
//}
//
//func ProcessChannel()  {
//
//}

//---------------------------------------------

func GenerateintA(done chan Student) chan int {
	ch := make(chan int, 10)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
	}()

	return ch
}

func GenerateintB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	return ch
}

//func GenerateInt() chan int {
//	ch := make(chan int,20)
//	go func() {
//		for  {
//			select {
//			case ch <- <- GenerateintA():
//			case ch <- <- GenerateintB():
//
//			}
//		}
//	}()
//	return ch
//}

func main() {
	done := make(chan Student)
	ch := GenerateintA(done)
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}
