package main

import "fmt"

func main() {
	//var conn *websocket.Conn
	//a := make(map[string]*websocket.Conn)
	//for _, v := range a {
	//	conn = v
	//	break
	//}
	//fmt.Println(conn)

	//
	//var addr = flag.String("addr", "localhost:8080", "http service address")
	//flag.Parse()
	//fmt.Println(*addr)

	c := make(chan int64)
	i, ok := <-c
	fmt.Println(i)
	fmt.Println(ok)
}
