package types

type User struct {
	Id   int
	name string
}

func (u *User) Name() string {
	return u.name
}
