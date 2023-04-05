package main

import "fmt"

const (
	name1 = "name"
	a     = iota
	b     = iota
)
const (
	name = "name"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
