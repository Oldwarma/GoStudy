package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	textStruct2json()
}

type Person2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func textStruct2json() {
	p := Person2{
		Name: "shaniao2",
		Age:  18,
	}
	fmt.Println("p", p)

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("person序列化失败")
		return
	}
	fmt.Println("json串", string(jsonBytes))
}
