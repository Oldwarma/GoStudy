package main

import (
	"fmt"
	"sync"
)

func main() {
	//n := flag.Int("n", 20, "Number of goroutines")
	//flag.Parse()
	//count := *n
	//fmt.Printf("Going to create %d goroutines.\n", count)
	var waitGroup sync.WaitGroup

	//fmt.Println("%#v\n", waitGroup)

	for i := 0; i < 20; i++ {
		waitGroup.Add(1) //确保每个协程都能运行
		go func(x int) {
			defer waitGroup.Done()
			fmt.Println(x, i)
		}(i)

	}
	waitGroup.Wait()
}
