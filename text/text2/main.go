package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6, 3, 78978}
	arr1 = append(arr1, arr2...)
	fmt.Println("arr1:", arr1)

	dataLength := int(binary.BigEndian.Uint32(buf))
}
