package main

import (
	"github.com/tealeg/xlsx"
	"log"
)

func main() {
	f, err := xlsx.OpenBinary("") // file为文件指针
	if err != nil {
		log.Fatal(err)
	}
}
