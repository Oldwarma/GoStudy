package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	if v := m["a"]; v != 0 { //B
		fmt.Println(v)
	}
}
