package main

import (
	"fmt"
	qr "github.com/skip2/go-qrcode"
	"github.com/valyala/fastjson"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func LoginWb() {
	qrid := QRCodeLogin1()

	QRCodeLogin2(qrid)
}

func QRCodeLogin1() string {
	getQRCodeUrl := "https://login.sina.com.cn/sso/qrcode/image?entry=weibo&size=180&callback=STK_16010457545441"
	header := make(map[string]string)
	header["referer"] = "https://weibo.com/"
	request, b := getRequest(getQRCodeUrl, header)
	if !b {
		return ""
	}
	split := strings.Split(string(request), "(")
	loginDataBytes := split[len(split)-1]
	loginDataBytes = strings.Trim(loginDataBytes, ");")
	value, err := fastjson.Parse(loginDataBytes)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(value.String())
	value = value.Get("data")
	qrid := string(value.GetStringBytes("qrid"))
	image := string(value.GetStringBytes("image"))
	fmt.Println(image)
	//exec.Command(`cmd`, `/c`, `start`, "https:"+image).Start()

	//出现二维码
	CreateQRCode("https:" + image)
	return qrid
}

func CreateQRCode(content string) string {
	fileName := "C:\\images\\" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg"
	err := qr.WriteFile(content, qr.Medium, 256, fileName)
	if err != nil {
		return ""
	}
	return fileName
}

func QRCodeLogin2(qrid string) {
	time.Sleep(time.Second * 20)
	header := make(map[string]string)
	header["referer"] = "https://weibo.com/"
	//进行检测
	getQRCodeUrl := "https://login.sina.com.cn/sso/qrcode/check?entry=weibo&qrid=" + qrid + "&callback=STK_16010457545443"
	bs, sucess := getRequest(getQRCodeUrl, header)
	if !sucess {

	}
	split := strings.Split(string(bs), "(")
	data := split[len(split)-1]
	data = strings.Trim(data, ");")
	value, err := fastjson.Parse(data)
	if err == nil {

	}
	retcode := value.GetInt64("retcode")
	switch retcode {
	case 20000000:
		data := value.Get("data")
		alt := string(data.GetStringBytes("alt"))
		getCookieApi := "https://login.sina.com.cn/sso/login.php?entry=weibo&returntype=TEXT&crossdomain=1&cdult=3&domain=weibo.com&alt=" + alt + "&savestate=30&callback=STK_160104719639113"
		response := HttpGetRequest(getCookieApi, nil)
		if !response.Sucess {
			fmt.Println("response失败")
			return
		}
		cookie := response.CookieString
		fmt.Println("普通cookie:", cookie)
		body := string(response.Data)
		fmt.Println("body:", body)
		compile := regexp.MustCompile("\\{[\\s\\S]*}")
		allString := compile.FindAllString(body, -1)
		fmt.Println("allString:", allString)
		parse, err := fastjson.Parse(allString[0])
		if err != nil {
			fmt.Println(err)
		}
		array := parse.GetArray("crossDomainUrlList")
		u := array[3].String()
		u = strings.Trim(u, `"`)
		u = strings.ReplaceAll(u, `\`, "")
		fmt.Println("访问链接u", u)
		finallyResponse := HttpGetRequest(u, nil)
		pcCookie := finallyResponse.CookieString
		getMbCKAPI := "https://login.sina.com.cn/sso/login.php?url=https%3A%2F%2Fm.weibo.cn%2F%3F%26jumpfrom%3Dweibocom&_rand=1588483688.7261&gateway=1&service=sinawap&entry=sinawap&useticket=1&returntype=META&sudaref=&_client_version=0.6.33"
		h := make(map[string]string)
		h["cookie"] = cookie
		mobileCookie := HttpGetRequest(getMbCKAPI, h)
		mbck := mobileCookie.CookieString
		fmt.Println("pcCookie:", pcCookie)
		fmt.Println("mbck:", mbck)

	//case 50114001:
	//case 50114003:
	//case 50114002:
	default:
		fmt.Println("蛤：", string(bs))

	}

	fmt.Println(string(bs))
}
