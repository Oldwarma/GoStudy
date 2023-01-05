package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const BufferSize = 1024 * 1024 * 300

func main() {
	for i := 1; i < 19; i++ {
		file, err := os.Open("./" + strconv.FormatInt(int64(i), 10) + ".txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		var j int64 = 1
		for {
			buffer := make([]byte, BufferSize)
			bs, err := file.Read(buffer)
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			WriteContent(strconv.FormatInt(int64(i), 10)+"-"+strconv.FormatInt(j, 10)+".txt", buffer[:bs])
			j++
		}
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
