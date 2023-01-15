package main

import (
	"flag"
	"fmt"
)

type Value interface {
	String() string
	Set(string) error
}

func main() {
	minusk := flag.Bool("k1212", true, "k")
	minusO := flag.Int("00", 1, "000")
	flag.Parse()
	valueK := *minusk
	valueO := *minusO
	valueO++
	fmt.Println("minusk", *minusk)
	fmt.Println("minusO", *minusO)

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)
}
