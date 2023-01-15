package main

import "fmt"

const (
	Zero digit = iota
	a
	b
)

type digit int

func main() {

	const c = iota
	const d = iota
	const e = iota
	fmt.Println(Zero, a, b, c, d, e)
}
