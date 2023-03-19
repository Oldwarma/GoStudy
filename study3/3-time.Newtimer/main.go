package main

import "time"

func main() {
	time.After(time.Second * 10)
	delay := time.NewTimer(time.Second * 10)
	defer delay.Stop()
	for {
		delay.Reset(time.Second * 10)
		select {
		case <-delay.C:

		}
	}
}

//使用 NewTimer 来做定时器，不需要每次都创建定时器对象。
//time.After 虽然调用的是timer定时器，但是他没有使用 time.Reset()方法再次激活定时器所以每一次都是新创建的实例，才会造成的内存泄漏。
