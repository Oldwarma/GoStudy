package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(random(1, 10))
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
