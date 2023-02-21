package main

import (
	"fmt"
	"gostudy/tools/utils"
	"io"
	"os"
	"path/filepath"
)

func main() {
	FilePath := "D:/cep/132112.tar.gz"
	FilePath, err := MkFile(FilePath)
	if err != nil {
		fmt.Println("err", err)
	}
}
func MkFile(filePath string) (resPath string, err error) {
	//取名
	fileName := filepath.Base(filePath)
	//取路径
	fileP := filepath.Dir(filePath)
	existPath := filepath.Join(fileP, "un")
	resPath = filepath.Join(fileP, "un", fileName)

	exist := utils.CheckFileExist(existPath) //检查配置目录是否存在
	if !exist {
		fmt.Println("confDir Create")
		utils.CreateDir(existPath)
	}

	dst, err := os.Create(resPath)
	if err != nil {
		fmt.Println("创建文件失败:", fileName)
		return "", err
	}
	defer dst.Close()
	src, err := os.Open(filePath)
	if err != nil {
		fmt.Println("复制文件失败:", fileName)
		return "", err
	}
	defer src.Close()
	//复制一份去解压
	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println("复制文件失败:", fileName)
		return "", err
	}
	return
}
