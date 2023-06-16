package main

import "fmt"

func main() {
	arr := make([]byte, 100)
	for i := 0; i < 100; i++ {
		arr[i] = byte(i)
	}
	for _, v := range arr {
		fmt.Println(v)
	}
}
