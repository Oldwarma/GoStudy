package main

import "fmt"

func main() {
	var myInt interface{} = 123
	k, ok := myInt.(int)
	if ok {
		fmt.Println("类型正确", k)
	} else {
		fmt.Println("类型错误", k)
	}
	v, ok := myInt.(float64)
	if ok {
		fmt.Println("类型正确", v)
	} else {
		fmt.Println("类型错误", v)
	}
	j, _ := myInt.(int32)
	fmt.Println(j)
}
