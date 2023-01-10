package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

type Person3 struct {
	Name    string `json:"name"`
	Aage    int    `json:"aage"`
	Address string `json:"address"`
}

func main() {
	textMap2Struct()
}
func textMap2Struct() {
	MapRes := make(map[string]interface{})
	MapRes["name"] = "shaniao5"
	MapRes["age"] = 18
	MapRes["address"] = "hongkong"

	var person Person3

	if err := gconv.Struct(MapRes, &person); err != nil {
		fmt.Println("map转struct失败")
		return
	}
	fmt.Println("map转换struct结果:", person)
}
