package main

import (
	"fmt"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}
type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	//Elem()
	Elem1()
}

func Elem() {
	x := "121"
	xRefl := reflect.ValueOf(&x).Elem()
	typexx := reflect.TypeOf(x) //直接取类型
	fmt.Printf("The type of x is %s.\n", typexx)

	for i := 0; i < xRefl.NumField(); i++ {
		fmt.Println("field name", xRefl.Type().Field(i).Name)
		fmt.Println("with type", xRefl.Type().Field(i))
		fmt.Println("Interface name", xRefl.Field(i).Interface())
	}
}

func Elem1() {
	A := a{100, 200.12, "struct a"}
	B := b{1, 2, "struct b", -1.2}
	fmt.Printf("A=%v\nB=%v", reflect.ValueOf(&A).Elem().Type().Name(), reflect.ValueOf(&B).Elem())
}
