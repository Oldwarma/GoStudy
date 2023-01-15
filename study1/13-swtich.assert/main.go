package main

import "fmt"

func main() {
	tellInterface(1)
	tellInterface(square{X: 4})
	tellInterface(circle{R: 10})
	tellInterface(rectangle{X: 4, Y: 1})
}

type square struct {
	X float64
}
type circle struct {
	R float64
}
type rectangle struct {
	X float64
	Y float64
}

func tellInterface(x interface{}) {
	switch x.(type) {
	case square:
		fmt.Println("square")
	case circle:
		fmt.Println("circle")
	case rectangle:
		fmt.Println("rectangle")
	default:
		fmt.Println("unknown type")
	}
}
