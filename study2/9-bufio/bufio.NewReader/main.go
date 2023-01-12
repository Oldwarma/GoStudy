package main

import (
	"archive/tar"
	"bufio"
	"fmt"
	"gostudy/tools/types"
	"io"
	"os"
)

func main() {
	newReader()
	tarNewReader()
}

func newReader() {
	file, err := os.Open(types.FilePath)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer file.Close()
	bufReader := bufio.NewReader(file) //io流，读取文件
	var buf [1024]byte
	for {
		n, err := bufReader.Read(buf[:])
		//bufReader.ReadLine()
		if err == io.EOF {
			fmt.Println("读取文件到头")
			fmt.Println("123", string(buf[:n]))
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("读取失败")
			break
		}
	}

}

func tarNewReader() {
	file, _ := os.Open("D:/cep/0110.tar")
	tr := tar.NewReader(file)

	for {
		hdr, err := tr.Next() //取压缩包里面的文件
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取结束")
			}
			break
		}
		fmt.Println(hdr.Name)
	}
	fmt.Println("ok")
}
