package main

import "fmt"

type arrs struct {
	age  string
	name string
}

func main() {
	var arr []arrs
	arr[0] = arrs{
		age: "56",
	}
	for _, v := range arr {
		if len(v.name) != 0 {
			fmt.Println("!=0")
			break
		} else {
			fmt.Println("==0")
			break
		}
		fmt.Println("0000")
	}
}
