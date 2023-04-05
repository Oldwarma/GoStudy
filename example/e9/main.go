package main

import "fmt"

func main() {
	var s1 []int
	var s2 = []int{}
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
	fmt.Printf("s1%v  s2%v", s1, s2)
}
