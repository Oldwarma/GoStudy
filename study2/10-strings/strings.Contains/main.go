package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	contains()
}

func contains() {
	ok := strings.Contains("1321321321/4654654/asdasd/", "*")
	fmt.Println(ok)
	select {
	case <-time.After(time.Second * 3):
		fmt.Println(ok, "time out")
	}
}
