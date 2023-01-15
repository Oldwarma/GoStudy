package main

import (
	"fmt"
	"sync"
)

var wg = new(sync.WaitGroup)

func main() {
	a := []string{"aaa", "bbb", "ccc"}
	for _, v := range a {
		fv := v
		wg.Add(1)
		go Run(fv)
	}
	wg.Wait()
	fmt.Println("ok")
}

func Run(id string) {
	defer wg.Done()
	var send = 0
	l := new(sync.Mutex)
	f := func() int {
		l.Lock()
		defer l.Unlock()
		send++
		return send
	}
	HandleFun(id, f)
	HandleFun(id, f)
	HandleFun(id, f)
}

type XXX func() int

func HandleFun(id string, seqFun XXX) {
	fmt.Println("ID为", id, "序列号为:", seqFun(), xxxxx())
}
func xxxxx() int {
	var send = 0

	send++
	return send
}
