package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cat(f *os.File) {
	i := 2
	for {
		buf := make([]byte, 302400000)
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			return
		case nr == 0: // EOF
			return
		case nr > 0:
			contains := strings.Contains(string(buf), "3608298---")
			index := strings.Index(string(buf), "3608298---")
			if index == -1 {
				continue
			}
			start := index - 10
			end := index + 20
			c := string(buf)
			content := c[start:end]
			if contains {
				writeFile("C:\\q\\"+strconv.FormatInt(int64(i), 10)+".txt", content)
				i++
				fmt.Println(content)
			}
		}
	}
}

func writeFile(fileName, content string) bool {
	fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(content)
	_, err := fd.Write(buf)
	fd.Close()
	if err == nil {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("D:\\Q\\q\\1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	cat(file)
	//fo, errs := os.Create(fmt.Sprintf("./%d.bmp", time.Now().UnixNano())) //time.Now().UnixNano()
	//fmt.Println(errs)
	//_, errs = fo.Write(payload)

}
