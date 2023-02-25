package types

import "fmt"

type User struct {
	Id       int
	NName    string
	nickName string
}

func (u *User) Name() string {
	return u.NName
}
func (u *User) NickName() string {
	return u.nickName
}

func (u *User) StringUser() {
	fmt.Println(u.Id, u.NName, u.nickName)
}
