package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testJson2map()
}
func testJson2map() {
	jsonStr := `
	{
		"name":"shaniao3",
		 "age":18
	}
	`
	var mapRes map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &mapRes)
	if err != nil {
		fmt.Println("json转map失败")
		return
	}
	fmt.Println("mapRes", mapRes)

}
