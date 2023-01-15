package main

import "fmt"

func main() {
	i := funReturnFun()
	j := funReturnFun()
	a := 0
	k := func() int {
		a++
		return a
	}
	fmt.Println("1", i())
	fmt.Println("2", i())
	fmt.Println("3", i())

	fmt.Println("j1", j())
	fmt.Println("j2", j())
	fmt.Println("j3", j())
	fmt.Println("4", i())

	fmt.Println("k1", k())
	fmt.Println("k2", k())
	fmt.Println("k3", k())
}

func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}
