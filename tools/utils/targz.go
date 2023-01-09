package utils

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func CompressDir(src, tarName string) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return errors.New("非文件夹")
	}
	outputFile, err := os.Create(tarName)
	gw := gzip.NewWriter(outputFile)
	defer gw.Close()

	tarWriter := tar.NewWriter(gw) // tarWriter:需要操作的句柄
	defer tarWriter.Close()

	dir, err := os.ReadDir(src)
	for _, v := range dir {
		info, err := os.Stat(path.Join(src, v.Name()))
		if err != nil {
			fmt.Println(111)
			fmt.Println(err)
			continue
		}
		if info.IsDir() {
			fmt.Println(222)
			fmt.Println("pass dir")
			continue
		}
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			fmt.Println(333)
			return err
		}
		stamp := filepath.Base(path.Join(src, v.Name()))
		header.Name = stamp
		err = tarWriter.WriteHeader(header)
		if err != nil {
			fmt.Println(444)
			return err
		}
		file, err := os.Open(path.Join(src, v.Name()))
		if err != nil {
			fmt.Println(555)
			return err
		}
		defer file.Close()
		buf := make([]byte, 4096)
		if _, err = io.CopyBuffer(tarWriter, file, buf); err != nil {
			fmt.Println(666)
			return err
		}
	}
	fmt.Println(777)
	return nil
}

func GenTarGz(filepath, filename string) error {
	File, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer File.Close()
	gw := gzip.NewWriter(File)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	return walk(filepath, "", tw)
}

func walk(dir, cDir string, tw *tar.Writer) error {
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, v := range info {
		if v.IsDir() {
			n := ""
			if cDir == "" {
				n = v.Name()
			} else {
				n = path.Join(cDir, v.Name())
			}
			head := tar.Header{Name: n, Typeflag: tar.TypeDir, ModTime: v.ModTime()}
			tw.WriteHeader(&head)
			walk(path.Join(dir, v.Name()), path.Join(cDir, v.Name()), tw)
			continue
		}
		F, err := os.Open(path.Join(dir, v.Name()))
		if err != nil {
			fmt.Print("打开文件失败:", err)
			continue
		}
		n2 := v.Name()
		if cDir != "" {
			n2 = path.Join(cDir, n2)
		}
		head := tar.Header{Name: n2, Size: v.Size(), Mode: int64(v.Mode()), ModTime: v.ModTime()}
		tw.WriteHeader(&head)
		io.Copy(tw, F)
		F.Close()
	}
	return nil
}
