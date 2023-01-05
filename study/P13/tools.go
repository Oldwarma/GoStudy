package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func downPic(u string) string {
	u = strings.ReplaceAll(u, "\"", "")
	request, _ := http.NewRequest("GET", u, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	//加入get参数
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()
	//isok := true
	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	defer func() {
		resp.Body.Close()
	}()
	if resp.Body == nil {
		return ""
	}
	body, _ := ioutil.ReadAll(resp.Body)
	name := `C:\images\` + strconv.FormatInt(time.Now().UnixNano(), 10) + `.png`
	out, _ := os.Create(name)
	defer out.Close()
	_, err = io.Copy(out, bytes.NewReader(body))
	if err != nil {
		return ""
	}
	return name
}

var (
	ckCompile = regexp.MustCompile(".*?;")
)

type Response struct {
	Sucess       bool
	Data         []byte
	Cookie       []*http.Cookie
	CookieString string
}

func HttpGetRequest(url string, headerData map[string]string) Response {

	var response Response

	if url == "" {
		return response
	}
	request, _ := http.NewRequest("GET", url, nil)
	if headerData != nil {
		for k, v := range headerData {
			request.Header.Set(k, v)
		}
	}
	//加入get参数
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()

	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		return response
	}
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return response
	}
	response.Data = data
	response.Cookie = resp.Cookies()
	defer func() {
		errs := resp.Body.Close()
		if errs != nil {
			fmt.Println(errs)
		}
	}()
	ck := ""
	values := resp.Header.Values("Set-Cookie")
	if values != nil {
		for _, v := range values {
			if strings.Contains(v, "deleted") || v == "" {
				continue
			}
			allString := ckCompile.FindAllString(v, -1)
			ck += allString[0] + " "
		}
	}
	if ck != "" {
		response.CookieString = ck
	}
	response.Sucess = true
	return response
}

func getRequest(url string, headerData map[string]string) ([]byte, bool) {
	if url == "" {
		return nil, false
	}
	request, _ := http.NewRequest("GET", url, nil)
	if headerData != nil {
		for k, v := range headerData {
			request.Header.Set(k, v)
		}
	}
	//加入get参数
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()
	isok := true

	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		isok = false
		return nil, false
	}

	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		isok = false
		return nil, false
	}
	defer func() {
		errs := resp.Body.Close()
		if errs != nil {
			fmt.Println(errs)
		}
	}()
	return data, isok
}

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func PostForm(api string, params, header map[string]string) ([]byte, []*http.Cookie) {
	data := make(url.Values)
	for k, v := range params {
		data[k] = []string{v}
	}

	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	reader := bytes.NewReader([]byte(data.Encode()))
	request, err := http.NewRequest("POST", api, reader)

	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}

	res, err := client.Do(request)
	cookies := res.Cookies()
	if err != nil {
		fmt.Println(err.Error())
		return nil, cookies
	}
	respBytes, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(respBytes))
	defer res.Body.Close()
	fmt.Println("post send success")
	return respBytes, cookies
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func CreateFormFile3(w *multipart.Writer, name, filename, contentType, contentLength string) (io.Writer, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(name), escapeQuotes(filename)))
	h.Set("Content-Type", contentType)
	h.Set("Content-Length", contentLength)
	return w.CreatePart(h)
}

func CreateFormFile4(w *multipart.Writer, name, contentType, contentLength string) (io.Writer, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"`, escapeQuotes(name)))
	if contentType != "" {
		h.Set("Content-Type", contentType)
	}
	if contentLength != "" {
		h.Set("Content-Length", contentLength)
	}
	return w.CreatePart(h)
}
