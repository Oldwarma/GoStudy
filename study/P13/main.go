package main

import (
	"fmt"
	"github.com/valyala/fastjson"
	"strings"
)

var (
	UpLoadPicRefer  = "https://m.weibo.cn/compose/"
	PublishWeiboApi = "https://m.weibo.cn/api/statuses/update"
	//PublishWeiboApi = "http://127.0.0.1:8886/test"
	PublishWeiboPrefixRefer = "https://m.weibo.cn/compose/?pids="

	PC_UA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36"
)

func main() {
	//mbck := "tgc=TGT-NzU4MDY5MzA0OA==-1618767985-gz-CDB89CA4D70F8170F6648F3A270D50D6-1; SUB=_2A25NeAAhDeRhGeFL41IX-S3MzzSIHXVugqBprDV_PUNbm9AfLVakkW1NfbA9AlFIJMUYNciCAvTYGUp26mVjI0w4; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9W5TzA765LsQO8.LQlEorh0B5NHD95QNSKn7So.0ehBRWs4DqcjMi--NiK.Xi-2Ri--ciKnRi-zNSo-41KeNeh5pSntt; ALC=ac%3D2%26bt%3D1618767985%26cv%3D5.0%26et%3D1650303985%26ic%3D1903579770%26login_time%3D1618767983%26scf%3D%26uid%3D7580693048%26vf%3D0%26vs%3D1%26vt%3D0%26es%3D315925c423fbf4874c683fcaa7e42476; ALF=1650303985; LT=1618767985;"

	moblieCookie := "tgc=TGT-NzU4MDY5MzA0OA==-1619018619-gz-5D9C0EADFFEFA63194064FAA8FC2EC8D-1; SUB=_2A25NhDMsDeRhGeFL41IX-S3MzzSIHXVuh11krDV_PUNbm9AfLWnakW1NfbA9AgmABDpisOmPxAFPycGDjaEOWl65; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9W5TzA765LsQO8.LQlEorh0B5NHD95QNSKn7So.0ehBRWs4DqcjMi--NiK.Xi-2Ri--ciKnRi-zNSo-41KeNeh5pSntt; ALC=ac%3D2%26bt%3D1619018620%26cv%3D5.0%26et%3D1650554620%26ic%3D1903579635%26login_time%3D1619018618%26scf%3D%26uid%3D7580693048%26vf%3D0%26vs%3D1%26vt%3D0%26es%3Df05aa132fac0f17c7aa00f1dbcd74944; ALF=1650554620; LT=1619018620;"
	weiboToken := GetWbToken(moblieCookie)
	//var ps = []string{"https://ss3.baidu.com/-fo3dSag_xI4khGko9WTAnF6hhy/image/h%3D300/sign=785454b5a41c8701c9b6b4e6177f9e6e/0d338744ebf81a4c07280698c02a6059252da64d.jpg"}
	publishWeibo(weiboToken, "大家好，我叫滑小稽123", nil)

}

func publishWeibo(wbToken WbToken, content string, picUrls []string) string {
	if wbToken.Cookie == "" && wbToken.Token == "" {
		fmt.Println("获取token失败")
		return ""
	}

	picIds := ""

	params := make(map[string]string)
	params["content"] = content
	params["st"] = wbToken.Token
	params["_spr"] = "screen:411x731"

	if len(picUrls) > 0 {
		for _, v := range picUrls {
			if v == "" {
				continue
			}
			id := UpLoadPic(v, "", UpLoadPicRefer, wbToken)
			fmt.Println(id)
			picIds += id + ","
		}
		picIds = strings.Trim(picIds, ",")
		params["picId"] = picIds
	}

	header := make(map[string]string)
	header["cookie"] = wbToken.Cookie
	header["Content-Type"] = "application/x-www-form-urlencoded"
	header["Connection"] = "Keep-Alive"
	header["referer"] = PublishWeiboPrefixRefer + picIds
	header["user-agent"] = PC_UA

	form, _ := PostForm(PublishWeiboApi, params, header)
	return string(form)
}

type WbToken struct {
	Token  string
	Cookie string
}

func GetWbToken(mbCookie string) WbToken {
	api := "https://m.weibo.cn/api/config"
	header := make(map[string]string)
	header["cookie"] = mbCookie
	request := HttpGetRequest(api, header)
	fmt.Println("获取token的respone:", string(request.Data))
	fmt.Println("获取token结果:", request.Sucess)
	parseBytes, err := fastjson.ParseBytes(request.Data)
	if err != nil {
		fmt.Println("解析token的respone错误:", err)
		return WbToken{}
	}
	data := parseBytes.Get("data")
	login := data.GetBool("login")
	if !login {
		fmt.Println("登录失效")
		return WbToken{}
	}
	token := string(data.GetStringBytes("st"))
	ck := request.CookieString + mbCookie

	return WbToken{token, ck}
}
