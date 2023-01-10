package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testMap2json()
}
func testMap2json() {
	MapRes := make(map[string]interface{})
	MapRes["name"] = "shaniao4"
	MapRes["age"] = 18
	MapRes["address"] = "北京"

	jsonByte, err := json.Marshal(MapRes)
	if err != nil {
		fmt.Println("map转json失败")
	}
	fmt.Println(string(jsonByte))
}
