package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var tempDir []string
	x := "C:\\Users\\Jiang\\AppData\\Local\\Temp"
	dir, _ := ioutil.ReadDir(x)
	for _, v := range dir {
		isDir := v.IsDir()
		if isDir {
			tempDir = append(tempDir, x+"\\"+v.Name())
		} else {
			if strings.Contains(v.Name(), ".jpg") || strings.Contains(v.Name(), "png") {
				err := os.Remove(x + "\\" + v.Name())
				fmt.Println(err == nil)
			}
		}
	}
	for _, v := range tempDir {
		dir, _ := ioutil.ReadDir(v)
		for _, v2 := range dir {
			if strings.Contains(v2.Name(), ".jpg") || strings.Contains(v2.Name(), "png") {
				err := os.Remove(v + "\\" + v2.Name())
				fmt.Println(err == nil)
			}
		}
	}
	wxPath := "C:\\Users\\Jiang\\Documents\\WeChat Files"
	wxPath2 := "\\FileStorage\\Image"
	wxPath3 := "\\FileStorage\\Image\\Thumb"
	var wxTempDir []string
	wxDir, _ := ioutil.ReadDir(wxPath)
	for _, v := range wxDir {
		imagePath := wxPath + "\\" + v.Name() + wxPath2
		_, err := os.Stat(imagePath)
		if err == nil {
			wxTempDir = append(wxTempDir, imagePath)
			imagePath = wxPath + "\\" + v.Name() + wxPath3
			_, err = os.Stat(imagePath)
			if err == nil {
				wxTempDir = append(wxTempDir, imagePath)
			}
		}
	}

	for _, v := range wxTempDir {
		fmt.Println(v)
		readDir, err := ioutil.ReadDir(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, v2 := range readDir {
			if err != nil {
				fmt.Println(err)
				continue
			}
			pathName := v + "\\" + v2.Name()
			files, err := ioutil.ReadDir(pathName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			for _, v3 := range files {
				if strings.Contains(v3.Name(), ".png") ||
					strings.Contains(v3.Name(), ".jpg") ||
					strings.Contains(v3.Name(), ".dat") {
					os.Remove(pathName + "\\" + v3.Name())
				}
			}
		}
	}
}
