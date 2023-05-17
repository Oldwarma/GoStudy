package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type user1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type user2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Text string `json:"text"`
}

func main() {
	temp1 := &user1{
		Name: "你干嘛",
		Age:  20,
	}
	var temp2 user2
	copier.Copy(&temp2, temp1)
	fmt.Println(temp2)
}
