package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	//HttpRequest()

	err := os.MkdirAll(`C:\Program Files (x86)\wind`, os.ModePerm)
	fmt.Println(err)
}

func HttpRequest() ([]byte, error) {

	data := make(map[string]string)

	req := &fasthttp.Request{} //相当于获取一个对象

	req.SetRequestURI("http://www.baidu.com") //设置请求的url

	bytes, err := json.Marshal(data) //data是请求数据

	if err != nil {
		return nil, err
	}

	req.SetBody(bytes) //存储转换好的数据

	req.Header.SetContentType("application/json") //设置header头信息

	req.Header.SetMethod("GET") //设置请求方法

	resp := &fasthttp.Response{} //相应结果的对象

	client := &fasthttp.Client{} //发起请求的对象

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	fmt.Println(string(resp.Body()))

	return nil, nil
}
