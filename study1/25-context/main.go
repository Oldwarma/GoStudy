package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	delay := 5
	//f1(delay)
	//f2(delay)
	f3(delay)
}
func f1(t int) {
	c1 := context.Background()
	c2, cancel := context.WithCancel(c1)
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		//cancel()
	}()
	select {
	case <-c1.Done():
		fmt.Println("f1():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1time():", r)
	}
	return
}
func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()
	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}
func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()
	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}
