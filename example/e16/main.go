package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]*Math{
	"foo": &Math{2, 3},
}

func main() {
	//tmp := m["foo"]
	//tmp.x = 4
	//m["foo"] = tmp
	//fmt.Println(m["foo"].x)
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
	fmt.Printf("%#v", m["foo"]) // %#v 格式化输出详细信息
}
