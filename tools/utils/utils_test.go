package utils

import (
	"fmt"
	"github.com/c4milo/unpackit"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	fileName := "D:\\nari-work\\云边融合_14752.rar"
	stat, _ := os.Stat(fileName)
	fmt.Println(stat.Size())
	now := time.Now()
	seek, _ := ReadFileWithSeek(fileName, 123, 231)
	end := time.Now()
	fmt.Println(len(seek))
	fmt.Println(end.Sub(now))
	now2 := time.Now()
	ReadFile(fileName)
	end2 := time.Now()
	fmt.Println(end2.Sub(now2))
}

func TestBytes2Str(t *testing.T) {
	str := Bytes2Str([]byte{})
	fmt.Println(str)
}

func TestUnTarGzPack(t *testing.T) {
	_, err := UnTarGzPack("D:/ep/11.tar.gz")
	fmt.Println(err)
}

func TestCompress(t *testing.T) {
	CompressOneFile("D:/cep/0-example", "D:\\cep\\11.tar.gz")
}

func TestGenerateArchiveFile(t *testing.T) {
	//CompressDir("D:/cep/agent","D:\\cep\\11.tar.gz")

	a := "D:/cep/agent"
	a = a[strings.LastIndex(a, "/")+1:]
	fmt.Println(a)
}

func TestFuc(t *testing.T) {
	var a int = 0
	var l = sync.Mutex{}
	f := func() int {
		l.Lock()
		defer l.Unlock()
		a++
		return a
	}
	f()
	f()
	f()
	fmt.Println(a)
}

func TestCopyFile(t *testing.T) {
	file, err := CopyFile("C:\\Users\\lin\\Desktop\\workDoc\\IEC8705.exe", "C:\\Users\\lin\\Desktop\\workDoc\\aaa\\sa.exe")
	fmt.Println(file)
	fmt.Println(err)
}

func TestUnTarGzPack2(t *testing.T) {
	pack, err := UnTarGzPack(`D:\nari-work\project\code\golang\cloud-edge-platform2.0\common\tools\utils\354`)
	fmt.Println(pack)
	fmt.Println(err)
}

func TestTarFolder(t *testing.T) {
	file, _ := os.Open(`D:\cep\trans\stat.tar.gz`)
	//tempDir := `D:\nari-work\cep-test\214-un`
	destPath, err := unpackit.Unpack(file, "D:\\cep\\trans\\un")
	fmt.Println(err)
	fmt.Println(destPath)
}

func TestCompressDir(t *testing.T) {
	//GenerateArchiveFile
	err := CompressDir("D:\\nari-work\\cep-test\\etc\\cert", "D:\\nari-work\\cep.tar.gz")
	fmt.Println(err)
}

func TestRm(t *testing.T) {
	err := os.RemoveAll("D:/EnglishFile")
	fmt.Println(err)
}

func TestUnTarGzPack222(t *testing.T) {
	pack2, err := UnTarGzPack2(`D:\nari-work\project\code\golang\cloud-edge-platform2.0\common\tools\utils\773`)
	fmt.Println(pack2)
	fmt.Println(err)
}
