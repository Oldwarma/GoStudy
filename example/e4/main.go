package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("aa")
	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				fmt.Println("1212")
			}
		}
	}()
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("aa")

		}
	}

}
