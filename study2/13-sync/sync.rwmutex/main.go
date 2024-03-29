package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

//如果你遇到可以明确区分 reader 和 writer goroutine 的场景，
//且有大量的并发读、少量的并发写，并且有强烈的性能需求，你就可以考虑使用读写锁 RWMutex 替换 Mutex。
func main() {
	var showFunction = func(c *secret) string { return "" }
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex!")
		showFunction = showWithLock
	}
	var waitGroup sync.WaitGroup
	fmt.Println("Pass:", showFunction(&Password))
	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass:", showFunction(&Password))
		}()
		go func() {
			waitGroup.Add(1)
			defer waitGroup.Done()
			Change(&Password, "123456")
		}()
		waitGroup.Wait()
		fmt.Println("Pass:", showFunction(&Password))
	}
}
func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}
func show(c *secret) string {
	c.RWM.RLock()
	fmt.Println("show")
	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}
func showWithLock(c *secret) string {
	c.M.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.M.Unlock()
	return c.password
}
