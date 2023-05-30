package main

import (
	"fmt"
	"time"
)

func main() {

	u := time.Now().UnixMilli()
	fmt.Println("uuu:", u)
}
