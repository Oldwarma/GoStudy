package st

var (
	ReadDev = make(chan int)
	ReadMap = make(map[int]chan int)
)
