package main

import (
	"bytes"
	"fmt"
	"github.com/valyala/fastjson"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

var (
	UPLOAD_PIC_API = "https://m.weibo.cn/api/statuses/uploadPic"
)

func UpLoadPic(picUrl, picName, referer string, wbToken WbToken) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("出错了:", err)
		}
	}()
	var err error
	var bytedata bytes.Buffer
	//转换成对应的格式
	Multipar := multipart.NewWriter(&bytedata)
	//添加您的镜像文件

	picLocal := ""
	if picName != "" {
		picLocal = picName
	} else {
		picLocal = downPic(picUrl)
	}
	bs, err := ioutil.ReadFile(picLocal)
	i := len(bs)
	formatInt := strconv.FormatInt(int64(i), 10)
	fileData, err := os.Open(picLocal)
	if err != nil {
		fmt.Println(err)
	}

	form, _ := CreateFormFile4(Multipar, "type", "", "4")
	form.Write([]byte("json"))

	form, err = CreateFormFile3(Multipar, "pic", "pic.jpg", "application/octet-stream", formatInt)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(form, fileData)

	form, _ = CreateFormFile4(Multipar, "st", "", "6")
	form.Write([]byte(wbToken.Token))

	form, _ = CreateFormFile4(Multipar, "_spr", "", "14")
	form.Write([]byte("screen:411x731"))

	Multipar.Close()
	req, err := http.NewRequest("POST", UPLOAD_PIC_API, &bytedata)

	// Don不要忘记设置内容类型,这将包含边界。
	req.Header.Set("Content-Type", Multipar.FormDataContentType())
	req.Header.Set("cookie", wbToken.Cookie)
	req.Header.Set("referer", referer)
	req.Header.Set("user-agent", PC_UA)
	fmt.Println(Multipar.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(req)
	defer func() {
		res.Body.Close()
	}()
	if err != nil {
		fmt.Println("提交错误", err)
		return ""
	}
	respBytes, err := ioutil.ReadAll(res.Body)

	parseBytes, err := fastjson.ParseBytes(respBytes)
	fmt.Println("提交返回数据", string(respBytes))
	pic_id := string(parseBytes.GetStringBytes("pic_id"))
	return pic_id
}
