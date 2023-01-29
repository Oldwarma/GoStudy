package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func main() {
	go monitor()
	var w sync.WaitGroup
	n := 10
	for r := 0; r < n; r++ {
		w.Add(1)
		go func() {
			defer w.Done()
			set(rand.Intn(10 * n))
		}()
	}
	w.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}
func set(newValue int) {
	writeValue <- newValue
}
func read() int {
	return <-readValue
}
func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d", value)
		case readValue <- value:
		}
	}
}
