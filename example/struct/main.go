package main

import "fmt"

type user struct {
	Name string
	Age  int
}

func main() {
	var person user
	person.Name = "ni"
	if person == (user{}) {
		fmt.Println("person为空")
	}

}
