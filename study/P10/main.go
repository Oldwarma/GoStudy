package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file := readFile("Eng.txt")
	split := strings.Split(file, "\r\n")
	for _, v := range split {
		split2 := strings.Split(v, " ")
		if len(split2) != 2 {
			continue
		}
		word := split2[0]
		property := split2[1]
		fmt.Println(word, property)
	}
}

func readFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return ""
	}
	return string(b)
}
