package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
)

func main() {
	for count := 1; count < 19; count++ {
		file, err := ioutil.ReadFile("./" + strconv.FormatInt(int64(count), 10) + ".txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		file = bytes.ReplaceAll(file, []byte("----"), []byte("-"))
		WriteContent("./"+strconv.FormatInt(int64(count), 10)+".txt", file)
		runtime.GC()
	}
}

func WriteContent(fileName string, content []byte) bool {
	fd, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	_, err := fd.Write(content)
	fd.Close()
	if err == nil {
		return true
	}
	return false
}
