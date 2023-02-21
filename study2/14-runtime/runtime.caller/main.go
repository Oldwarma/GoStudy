package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("getCurrentAbPathByCaller = ", getCurrentAbPathByCaller())

}

//这是由于go run会将源代码编译到系统TEMP或TMP环境变量目录中并启动执行；
//而go build只会在当前目录编译出可执行文件，并不会自动执行

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// 获取当前执行程序所在的绝对路径 (go run )--
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}
