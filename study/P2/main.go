package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	spFlag := "-"
	for count := 1; count < 19; count++ {
		for j := 1; j < 5; j++ {
			fileName := strconv.FormatInt(int64(count), 10) + "-" + strconv.FormatInt(int64(j), 10) + ".txt"
			ok := CheckFileExist(fileName)
			if !ok {
				continue
			}
			file, err := ioutil.ReadFile("./" + fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			//多层文件夹名字--data
			dirMap := make(map[string]map[string]string)
			//infoMap := make(map[string]string)
			fmt.Println("开始创建文件")
			split := strings.Split(string(file), "\r\n")
			for k, v := range split {
				if strings.Contains(v, spFlag) {
					if k == 0 || k == len(split)-1 {
						fmt.Println("第", count, "-", j, "个文件:", v)
					}
					split2 := strings.Split(v, spFlag)
					QQ := split2[0]
					QQRunes := []rune(QQ)
					if len(QQRunes) < 5 {
						fmt.Println("第", count, "-", j, "个文件:", v)
						continue
					}
					Phone := split2[1]
					//文件夹名
					dirName := string(QQRunes[:4])
					QQ = string(QQRunes[4:])
					//infoMap[QQ] = Phone
					data, ok := dirMap[dirName]
					if ok {
						data[QQ] = Phone
						dirMap[dirName] = data
					} else {
						data = make(map[string]string)
						data[QQ] = Phone
						dirMap[dirName] = data
					}
				} else {
					if k == 0 {
						fmt.Println("第", count, "-", j, "个文件:", v)
					}
				}
			}
			fmt.Println("第", count, "-", j, "次GC")
			runtime.GC()
			//k路径，v数据
			for k, v := range dirMap {
				ok := CheckFileExist(k)
				if !ok {
					_ = os.MkdirAll("./"+k, 0777)
				}
				fn := k + "/" + strconv.FormatInt(int64(count), 10) + "-" + strconv.FormatInt(int64(j), 10) + "-data.txt"
				encode, err := Encode(v)
				if err != nil {
					fmt.Println(err)
					return
				}
				WriteContent("./"+fn, encode)
				runtime.GC()
			}
			fmt.Println("蛤")
			runtime.GC()
		}
	}
}

// 用gob进行数据编码
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 用gob进行数据解码
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
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

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
