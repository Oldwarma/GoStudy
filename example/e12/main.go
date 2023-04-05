package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age, "1")

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age, "2")
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age, "3")
	}()

	//person.age = 29
	person = &Person{29}
}
