package utils

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/c4milo/unpackit"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFile(filePath string) []byte {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	return file
}

func WriteFile(fileName string, content []byte) bool {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	defer func() {
		f.Close()
	}()
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		write, err := f.Write(content)
		if err == nil && write > 0 {
			return true
		}
	}
	return false
}

func WriteArrayFile(fileName string, contents []string) bool {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	defer func() {
		f.Close()
	}()
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		for _, content := range contents {
			write, err := f.Write([]byte(content))
			if err != nil && write == 0 {
				return false
			}
		}
	}
	return true
}

func ReadFileWithSeek(filePath string, seek, max int64) ([]byte, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return nil, errors.New("文件不存在")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("打开文件失败")
	}
	defer file.Close()
	file.Seek(seek, 0)
	bufReader := bufio.NewReader(file)
	fileSize := 0
	if stat.Size()-seek > max {
		fileSize = int(max)
	} else {
		fileSize = int(stat.Size() - seek)
	}
	buf := make([]byte, fileSize)
	_, err = bufReader.Read(buf)
	if err != nil && err != io.EOF {
		return nil, errors.New("文件读取失败")
	}
	return buf, nil
}

func CheckFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}

func CreateDir(filePath string) bool {
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		fmt.Println("create error:", err)
		return false
	}
	return true
}

//追加数据
func AppendFile(fileName string, content []byte) bool {
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//content = append(content, 0x0A)
	_, err := f.Write(content)
	f.Close()
	if err == nil {
		return true
	}
	return false
}

//删除数据
func DeleteFile(filePath string, fileName ...string) {
	for _, v := range fileName {
		join := path.Join(filePath, v)
		err := os.Remove(join)
		if err != nil {
			fmt.Println("删除失败", err)
		}
	}
}

func DeleteFileV2(filePath string) {
	err := os.Remove(filePath)
	fmt.Println(err)
}

func DeleteOneFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		return
	}
}

//解压.tar.gz
func UnTarGzPack(fileName string) (string, error) {
	srcFile, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return "", err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	//取文件路径
	dir := ""
	if strings.Contains(fileName, `/`) {
		dir = string([]rune(fileName)[0:strings.LastIndex(fileName, "/")])
	} else if strings.Contains(fileName, `\`) {
		dir = string([]rune(fileName)[0:strings.LastIndex(fileName, `\`)])
	}
	//取文件名
	newDir := ""
	if strings.Contains(fileName, `/`) {
		newDir = string([]rune(fileName)[strings.LastIndex(fileName, "/")+1:])
	} else if strings.Contains(fileName, `\`) {
		newDir = string([]rune(fileName)[strings.LastIndex(fileName, `\`)+1:])
	}
	if strings.Contains(fileName, ".") {
		newDir = string([]rune(newDir)[:strings.Index(newDir, ".")])
	}
	//变成文件路径+文件名的文件夹
	dir = dir + "/UnTarGzPack/" + newDir + ""
	exist := CheckFileExist(dir)
	if !exist {
		CreateDir(dir)
	}
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		}
		if hdr.FileInfo().IsDir() {
			filename := dir + "/" + hdr.Name
			//CreateDir(fileName)

			readDir, err := ioutil.ReadDir(filename)
			fmt.Println(err)
			fmt.Println(readDir)
		} else {
			filename := dir + "/" + hdr.Name
			file, err := createFile(filename)
			if err != nil {
				return "", err
			}
			io.Copy(file, tr)
		}
	}
	return dir, nil
}

func createFile(name string) (*os.File, error) {
	//ok := CreateDir(string([]rune(name)[0:strings.LastIndex(name, "/")]))
	//if ok {
	//	return nil, errors.New("创建文件夹失败:"+name)
	//}
	return os.Create(name)
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
	return nil
}

func Compress(files *os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	err := compress(files, "", tw)
	if err != nil {
		return err
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func CompressOneFile(src, targetName string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("找不到文件:", err)
		return err
	}
	header, err := tar.FileInfoHeader(srcInfo, "")
	if err != nil {
		fmt.Println(src, "6666666666666")
		return err
	}
	header.Name = srcInfo.Name()
	file, err := os.Open(src)
	fmt.Println(src, "000000")
	if err != nil {
		return err
	}
	defer file.Close()

	outputFile, err := os.Create(targetName)
	gw := gzip.NewWriter(outputFile)
	defer gw.Close()

	tarWriter := tar.NewWriter(gw) // tarWriter:需要操作的句柄
	defer tarWriter.Close()

	err = tarWriter.WriteHeader(header)
	fmt.Println(src, targetName, "1111")
	if err != nil {
		fmt.Println(err)
		return err
	}
	buf := make([]byte, 4096)
	if _, err = io.CopyBuffer(tarWriter, file, buf); err != nil {
		return err
	}
	return nil
}

func GetCurrentAbsPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

//只能传入绝对路径
func CopyFile(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	if strings.Contains(destFile, `\`) || strings.Contains(destFile, `/`) {
		dir := ""
		if runtime.GOOS == "windows" {
			dir = destFile[:strings.LastIndex(destFile, `\`)]
		} else {
			dir = destFile[:strings.LastIndex(destFile, `/`)]
		}
		if !CheckFileExist(dir) {
			CreateDir(dir)
		}
	}

	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	// 拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			return total, nil
		} else if err != nil {
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
}

func UnTarGzPack2(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("解压文件失败:", fileName)
		return "", err
	}
	undir := ""
	if runtime.GOOS == "windows" {
		if strings.Contains(fileName, `\`) {
			undir = path.Join(fileName[:strings.LastIndex(fileName, `\`)])
		}
	} else {
		if strings.Contains(fileName, `/`) {
			undir = path.Join(fileName[:strings.LastIndex(fileName, `/`)])
		}
	}
	fmt.Println("解压目录：", undir)
	_, err = unpackit.Unpack(file, path.Join(undir, "un"))
	return path.Join(undir, "un"), err
}
