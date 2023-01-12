package main

import (
	"fmt"
	"strings"
)

func main() {
	lastIndex()
}

//最后一个符号的位置
func lastIndex() {
	n := strings.LastIndex("31321321/13213/1adasd/13213/", "/")
	fmt.Println(n)
}
