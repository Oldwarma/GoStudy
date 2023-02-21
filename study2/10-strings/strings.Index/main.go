package main

import (
	"fmt"
	"strings"
)

var (
	str = "21321321/31321/31321/"
)

func main() {
	index()
	fmt.Println(index2("/123/456/123456/", "/"))
}

func index() {
	res := strings.Index(str, "/")
	fmt.Println(str[:res])
}

func index2(s, substr string) int {
	n := len(substr)
	switch {
	case n == 0:
		return 0
	case n == len(s):
		if substr == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	return 10086
}
