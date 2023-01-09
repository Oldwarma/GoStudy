package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

type HttpResp struct {
	Err       error
	Data      []byte
	Localtion string
}

var (
	httpTimeout = 2000 * time.Second
	httpClient  = http.Client{
		Timeout: httpTimeout,
	}
)

func HttpPostWithJson(api string, param interface{}) HttpResp {
	var result HttpResp
	defer func() {
		e := recover()
		if e != nil {
			logx.Error(e)
		}
	}()
	var bytesData []byte
	if param != nil {
		bytesData, _ = json.Marshal(param)
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", api, reader)
	if err != nil {
		result.Err = err
		return result
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := httpClient.Do(request)
	if err != nil {
		result.Err = err
		return result
	}
	if resp == nil {
		result.Err = errors.New("空数据")
		return result
	}
	defer func() {
		resp.Body.Close()
	}()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.Err = err
		return result
	}
	result.Data = respBytes
	return result
}

func HttpGet(api string) HttpResp {
	var result HttpResp
	request, _ := http.NewRequest("GET", api, nil)
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()
	resp, err := httpClient.Do(request)
	if err != nil {
		result.Err = err
		fmt.Println("下载失败:", err)
		return result
	}
	defer func() {
		resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	result.Data = body
	return result
}

func HttpDown(api, fileName string) bool {
	result := HttpGet(api)
	dir := ""
	if strings.Contains(fileName, `\`) {
		sp := strings.Split(fileName, `\`)
		dir = sp[len(sp)-1]
	} else if strings.Contains(fileName, `/`) {
		sp := strings.Split(fileName, `/`)
		dir = sp[len(sp)-1]
	}
	if !CheckFileExist(dir) {
		CreateDir(dir)
	}
	out, _ := os.Create(fileName)
	defer out.Close()
	_, err := io.Copy(out, bytes.NewReader(result.Data))
	return err == nil
}

func PostMultipartImage(filePath, url string) HttpResp {
	var resp HttpResp
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(PrintStackTrace(err))
		}
	}()
	name := filePath
	if strings.Contains(filePath, `\`) {
		name = filePath[strings.LastIndex(filePath, `\`)+1:]
	} else if strings.Contains(filePath, `/`) {
		name = filePath[strings.LastIndex(filePath, `/`)+1:]
	}
	var bytedata bytes.Buffer
	//转换成对应的格式
	Multipar := multipart.NewWriter(&bytedata)
	//添加您的镜像文件
	fileData, err := os.Open(filePath)
	if err != nil {
		resp.Err = err
		return resp
	}
	defer fileData.Close()

	//这里添加图片数据
	form, err := Multipar.CreateFormFile("file", name)
	if err != nil {
		resp.Err = err
		return resp
	}
	if _, err = io.Copy(form, fileData); err != nil {
		resp.Err = err
		return resp
	}
	Multipar.Close()

	//现在你有一个表单,你可以提交它给你的处理程序。
	req, err := http.NewRequest("POST", url, &bytedata)
	if err != nil {
		resp.Err = err
		return resp
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Don不要忘记设置内容类型,这将包含边界。
	req.Header.Set("Content-Type", Multipar.FormDataContentType())
	//提交请求

	res, err := httpClient.Do(req)
	if err != nil {
		resp.Err = err
		return resp
	}
	respBytes, err := ioutil.ReadAll(res.Body)
	resp.Data = respBytes
	return resp
}
