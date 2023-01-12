package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//readDir()

	DoGzip([]string{"hello，golang"}, "/cep/aaa.gz", "test.txt")
}

func readDir() {
	files, _ := ioutil.ReadDir("D:/cep/")
	for _, v := range files {
		fmt.Println(v.Name())
	}
}

/**
压缩bytes内容
1.根据指定目录创建文件
2.根据文件资源对象生成gzip Writer对象
3.往gzip Writer对象写入内容
4.err := doGzip([]byte("hello，golang"),"/data/aaa.gz","test.txt")
*/
func DoGzip(contents []string, path string, fileName string) error {
	gzFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer gzFile.Close()
	gzipWriter := gzip.NewWriter(gzFile)

	defer gzipWriter.Close()
	gzipWriter.Name = fileName
	for _, content := range contents {
		_, err = gzipWriter.Write([]byte(content))
		if err != nil {
			return err
		}
	}
	fmt.Println("结束")
	return nil
}
