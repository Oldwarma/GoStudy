package main

import (
	"encoding/json"
	"fmt"
)

// Person 结构体字段首字母必须要大写，否则外部无法访问，json转struct也转不成功
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Per struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func main() {
	testJson2Struct()
	p1 := Per{
		id:   '2',
		name: "sss",
	}
	fmt.Println("p1", p1)
	fmt.Println("&p1", &p1)
}

func testJson2Struct() {
	jsonStr := `
{
      "name":"shaniao",
       "age":"23"
}
`
	var person Person
	fmt.Println(&person, person)
	json.Unmarshal([]byte(jsonStr), &person)
	fmt.Println(&person, person)
	fmt.Println("person:", person)
}
