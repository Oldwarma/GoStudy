package lib

type Ter interface {
	Send(msg string) error
	Code() error
	Moniter() (string, error)
}
