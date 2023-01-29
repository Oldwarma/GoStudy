package main

import (
	"fmt"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}
type Data struct {
	job    Client
	square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func main() {
	nJobs := 10
	nWorkers := 5
	go create(nJobs)
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Println("Client ID: %d\tint: ", d.job.id)
			fmt.Println("%dtsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()
	makeWp(nWorkers)
	fmt.Println(": %v\n", <-finished)
}

func worker(w *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}
func makeWp(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}
func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}
