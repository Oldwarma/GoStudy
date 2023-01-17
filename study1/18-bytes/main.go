package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	var buffer bytes.Buffer //缓存流
	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n")
	buffer.WriteTo(os.Stdout)
	buffer.WriteTo(os.Stdout)
	//Reset() 方法重置 buffer 变量，通过 write() 方法再次写入一些数据。
	//然后你可以通过 bytes.NewReader() 创建新的读对象，再使用 io.Reader 接口方法 Read() 从 buffer 变量读取数据。
	buffer.Reset()
	buffer.Write([]byte("Mastering Go!"))
	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())
	for {
		b := make([]byte, 4)
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}
}
