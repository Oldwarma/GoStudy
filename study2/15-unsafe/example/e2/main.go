package main

import (
	"gostudy/study2/15-unsafe/types"
	"unsafe"
)

func main() {
	u := types.User{
		Id:    1,
		NName: "2",
	}
	//println("u.name:", u.Name())
	//println("unsafe.Pointer(&u)", unsafe.Pointer(&u))
	//println("unsafe.Sizeof(u.Id)", unsafe.Sizeof(u.Id))
	//n := (*string)(unsafe.Add(unsafe.Pointer(&u), unsafe.Sizeof(u.Id)))
	//*n = "modeify by unsafe"
	//println("u.name:", u.Name())

	//println("unsafe.Pointer(&u)", unsafe.Pointer(&u))
	println("unsafe.Sizeof(u.Id)", unsafe.Sizeof(u.NName))
	s := (*string)(unsafe.Add(unsafe.Pointer(&u), unsafe.Sizeof(u.NName)+unsafe.Sizeof(u.Id)))
	*s = "modeify by unsafe"
	u.StringUser()
}
