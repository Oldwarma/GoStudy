package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	parseDate()
	fmt.Println("ok")
}
func unix() {
	t := time.Now().Unix() //时间戳
	t1 := time.Now()
	fmt.Println(t, t1)
}
func parseDate() {
	var myDate string
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s date\n",
			filepath.Base(os.Args[0]))
		os.Exit(0)
	}
	myDate = os.Args[1]
	d, err := time.Parse("02 January 2006", myDate)
	if err == nil {
		fmt.Println("Full", d)
		fmt.Println("Time", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}
}
