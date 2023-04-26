package main

import "fmt"

func main() {
	b := byte(0x56)
	mask := byte(1 << 2)
	value := (b & mask) >> 2
	fmt.Println(value)

	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}
}
