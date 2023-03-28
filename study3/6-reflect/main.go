package main

import (
	"fmt"
	"reflect"
)

func main() {
	//worker1()
	worker2()
	worker3()
	for {

	}
}

type Service struct {
	Name string `json:"name2"`
}

func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) SetName(name string) {
	s.Name = name
}

func worker1() {
	s := Service{}
	s.SetName("chihuo")
	name := s.GetName()
	fmt.Printf("name1:= %s\n", name)
}

func worker2() {
	s := Service{}
	rv := reflect.ValueOf(&s)
	params := []reflect.Value{reflect.ValueOf("chihuo2")}
	rv.MethodByName("SetName").Call(params)

	ret := rv.MethodByName("GetName").Call(nil)

	fmt.Printf("reflect name2:%s\n", ret[0].String())
}

func worker3() {
	s := Service{}
	rt := reflect.TypeOf(s)
	if field, ok := rt.FieldByName("Name"); ok {
		tag := field.Tag.Get("json")
		fmt.Printf("filed tag is %s\n", tag)
	}
}
