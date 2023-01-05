package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/valyala/fastjson"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	a                  = false
	DATANAME2   string = "C:\\winvcl\\"
	DATAFOLDER2 string = "C:\\winvcl"
	IV                 = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	KEY                = "dat=@vcl"
	SK                 = DATANAME2 + "sk.json"
	//oauthUrl string = "http://159.75.9.136:8800/oauth"
	oauthUrl string = "http://127.0.0.1:8800/oauth"
	RobotMap        = make(map[int64]string)
)

func checkFile(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func main() {
	checkQQAuth()
	fmt.Println(a)
}

func checkQQAuth() {
	RobotMap[123] = "123"
	//校验本地文件
	exist := checkFile(SK)
	//如果无，生成3月1过期的
	if !exist {
		//unix := time.Now().Unix()
		ot := "1614622140"
		encrypt, _ := Encrypt([]byte(ot), []byte(KEY), IV)
		encoded := base64.StdEncoding.EncodeToString(encrypt)
		ok := checkFile(DATAFOLDER2)
		if !ok {
			_ = os.Mkdir(DATAFOLDER2, 0777)
		}
		WriteFile(DATANAME2+"sk.json", encoded)
	}
	//如果有，校验结束时间
	content := readFile(SK)

	if content == "" {
		a = true
		return
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		a = true
		return
	}
	decrypt, err := Decrypt(decodeBytes, []byte(KEY), IV)
	if err != nil {
		a = true
		return
	}
	parseInt, _ := strconv.ParseInt(string(decrypt), 10, 64)
	//超时
	if parseInt < time.Now().Unix() {
		data := make(map[string]string)
		QQ := ""
		for k, _ := range RobotMap {
			QQ = strconv.FormatInt(k, 10) + "|"
		}
		QQ = strings.Trim(QQ, "|")
		data["lala"] = QQ
		isOk, result := requestPost(data, oauthUrl)
		if !isOk {
			a = true
		} else {
			parse, err := fastjson.Parse(result)
			if err != nil {
				a = true
			} else {
				getInt64 := parse.GetInt64("code")
				if getInt64 == 0 {
					ot := time.Now().Unix() + 259200
					ots := strconv.FormatInt(ot, 10)
					encrypt, _ := Encrypt([]byte(ots), []byte(KEY), IV)
					encoded := base64.StdEncoding.EncodeToString(encrypt)
					ok := checkFile(DATAFOLDER2)
					if !ok {
						_ = os.Mkdir(DATAFOLDER2, 0777)
					}
					WriteFile(DATANAME2+"sk.json", encoded)
					a = true
					return
				} else {
					a = true
					return
				}
			}
		}
	} else {
		//未超时
		a = true
		return
	}
}

func Encrypt(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func Decrypt(crypted, key []byte, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func readFile(fileName string) string {
	if !checkFile(fileName) {
		return ""
	}
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return ""
	}
	return string(b)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//复写
func WriteFile(fileName, content string) bool {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer func() {
		f.Close()
	}()
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		write, err := f.Write([]byte(content))
		if err == nil && write > 0 {
			return true
		}
	}
	return false
}

func requestPost(data interface{}, url string) (bool, string) {
	bytesData, _ := json.Marshal(data)
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return false, "err"
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	if err != nil {
		return false, "err"
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "err"
	}

	content := string(respBytes)
	return true, content
}
