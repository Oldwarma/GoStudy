package main

import (
	"fmt"
	"time"
)

func myFunction() {
	fmt.Println("Hello world!")
}

func main() {

	go func() {
		pro()
	}()
	for {

	}
}

func pro() {
	delay := time.NewTimer(time.Second * 5)
	defer delay.Stop()
	for {
		delay.Reset(time.Second * 5)
		select {
		case <-delay.C:
			myFunction()
		}
	}
}
