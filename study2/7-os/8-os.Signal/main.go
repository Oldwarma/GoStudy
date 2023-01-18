package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle(signal os.Signal) {
	fmt.Println("Received:", signal)
}
func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		sig := <-sigs
		switch sig {
		case os.Interrupt:
			handle(sig)
		case syscall.SIGTERM:
			handle(sig)
			os.Exit(0)
		case syscall.SIGINT:
			fmt.Println("Handling syscall.SIGUSR2!")
		default:
			fmt.Println("Ignoring:", sig)
		}

	}()
	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}
}
