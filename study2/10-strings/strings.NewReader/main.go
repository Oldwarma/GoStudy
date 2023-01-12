package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	newReader()
}
func newReader() {
	//file, err := os.Open(types.FilePath)
	//if err != nil {
	//	fmt.Println("打开文件失败")
	//}
	//defer file.Close()
	//bufReader := bufio.NewReader(file) //io流，读取文件
	file := strings.NewReader("testLine")
	var buf [1024]byte
	for {
		_, err := file.Read(buf[:])
		if err == io.EOF {
			fmt.Println("读取完成", string(buf[:8]))
			break
		}
	}

}
