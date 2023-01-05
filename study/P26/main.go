package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t.Unix())
		}
	}
}
