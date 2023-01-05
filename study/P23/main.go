package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name string
	Age  int64
}

func Generation() chan Student {
	c := make(chan Student, 1)
	go func() {
		var s Student
		time.Sleep(time.Second * 1)
		s.Name = "滑稽"
		fmt.Println("返回数据")
		c <- s
		close(c)
	}()
	return c
}

func main() {
	c := Generation()
	var r Student
	select {
	case r = <-c:
		fmt.Println("收到学生信息")
		break
	case <-time.After(time.Second * 3):
		fmt.Println("超时")
		break
	}
	fmt.Println(r)
	time.Sleep(time.Second * 8)
}
