package main

import "gostudy/study3/1-codingfaster/lib"

type Msg struct {
}

func (m *Msg) Send(msg string) error {
	//TODO implement me
	panic("implement me")
}

func (m *Msg) Code() error {
	//TODO implement me
	panic("implement me")
}

func (m *Msg) Moniter() (string, error) {
	//TODO implement me
	panic("implement me")
}

func xx() {
	msg := &Msg{}
	SessMsg(msg)

}

func SessMsg(req lib.Ter) {

}
