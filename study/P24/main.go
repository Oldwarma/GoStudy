package main

import (
	"fmt"
	"regexp"
	"strings"
)

var html = `
  <div align="center" id="content">
   <div align="center" class="mdui-typo-title">
    第7周 课程表(2021-2022第1学期)
   </div>
   <div align="center" class="f3">
    打印时间：2021-10-15 21:16:56
   </div>
   <table align="center" border="0" cellpadding="0" cellspacing="0" class="mdui-table  mdui-table-hoverable">
    <tbody>
     <tr height="30">
      <td align="center" class="t2 l2 b1 r1 b" width="60">
      </td>
      <td align="center" class="t2 b1 r1 b" width="100">
       星期一
       <br/>
       2021-10-18
      </td>
      <td align="center" class="t2 b1 r1 b" width="100">
       星期二
       <br/>
       2021-10-19
      </td>
      <td align="center" class="t2 b1 r1 b" width="100">
       星期三
       <br/>
       2021-10-20
      </td>
      <td align="center" class="t2 b1 r1 b" width="100">
       星期四
       <br/>
       2021-10-21
      </td>
      <td align="center" class="t2 b1 r1 b" width="100">
       星期五
       <br/>
       2021-10-22
      </td>
      <td align="center" class="t2 b1 r1 b" width="100">
       星期六
       <br/>
       2021-10-23
      </td>
      <td align="center" class="t2 b1 r2 b" width="100">
       星期日
       <br/>
       2021-10-24
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第1节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="2">
       <div title="职业生涯规划
卢元文
107(本部实验楼)">
        <div class="mdui-typo-subheading">
         职业生涯规划
        </div>
        <div>
         卢元文
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          107(本部实验楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         8:20-09:40
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="2">
       <div title="网络设备配置与调试
赵飞
305(本部特教大楼)">
        <div class="mdui-typo-subheading">
         网络设备配置与调试
        </div>
        <div>
         赵飞
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          305(本部特教大楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         08:40-10:00
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="2">
       <div title="C语言程序设计
刘永志
401(本部特教大楼)">
        <div class="mdui-typo-subheading">
         C语言程序设计
        </div>
        <div>
         刘永志
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          401(本部特教大楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         8:20-09:40
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第2节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第3节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="2">
       <div title="网络设备配置与调试
赵飞
401(本部特教大楼)">
        <div class="mdui-typo-subheading">
         网络设备配置与调试
        </div>
        <div>
         赵飞
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          401(本部特教大楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         10:10-11:30
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="2">
       <div title="C语言程序设计
刘永志
305(本部特教大楼)">
        <div class="mdui-typo-subheading">
         C语言程序设计
        </div>
        <div>
         刘永志
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          305(本部特教大楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         10:30-11:50
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="2">
       <div title="职业外语1
张旭丹
115(本部综合楼)">
        <div class="mdui-typo-subheading">
         职业外语1
        </div>
        <div>
         张旭丹
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          115(本部综合楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         10:10-11:30
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第4节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       中午1
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       中午2
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第5节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="2">
       <div title="数据库应用与安全
吴祥容
401(本部特教大楼)">
        <div class="mdui-typo-subheading">
         数据库应用与安全
        </div>
        <div>
         吴祥容
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          401(本部特教大楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         13:50-15:10
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="3">
       <div title="思想道德与法治（理论）
吴晓芳
实验楼411(本部实验楼)">
        <div class="mdui-typo-subheading">
         思想道德与法治（理论）
        </div>
        <div>
         吴晓芳
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          实验楼411(本部实验楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         13:50-16:20
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="2">
       <div title="数据库应用与安全
吴祥容
305(本部特教大楼)">
        <div class="mdui-typo-subheading">
         数据库应用与安全
        </div>
        <div>
         吴祥容
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
          305(本部特教大楼)
         </font>
        </div>
       </div>
       <font>
        <b>
         14:10-15:30
        </b>
       </font>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="3">
       <div title="思想道德与法治（实践）
吴晓芳
">
        <div class="mdui-typo-subheading">
         思想道德与法治（实践）
        </div>
        <div>
         吴晓芳
        </div>
        <div class="mdui-typo-subheading">
         <font color="87ACF2">
         </font>
        </div>
       </div>
       <font>
        <b>
         14:00-16:30
        </b>
       </font>
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第6节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第7节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第8节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第9节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第10节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b1 r1 b">
       第11节
       <br/>
       <br/>
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r1" rowspan="1">
      </td>
      <td align="center" class="b1 r2" rowspan="1">
      </td>
     </tr>
     <tr>
      <td align="center" class="l2 b2 r1 b">
       第12节
       <br/>
       <br/>
      </td>
      <td align="center" class="b2 r1" rowspan="1">
      </td>
      <td align="center" class="b2 r1" rowspan="1">
      </td>
      <td align="center" class="b2 r1" rowspan="1">
      </td>
      <td align="center" class="b2 r1" rowspan="1">
      </td>
      <td align="center" class="b2 r1" rowspan="1">
      </td>
      <td align="center" class="b2 r1" rowspan="1">
      </td>
      <td align="center" class="b2 r2" rowspan="1">
      </td>
     </tr>
    </tbody>
   </table>
  </div>
 </body>
</html>`

func main() {
	html = strings.ReplaceAll(html, "\r\n", "")
	html = strings.ReplaceAll(html, "\r", "")
	html = strings.ReplaceAll(html, "\n", "")
	trReg := regexp.MustCompile(`<tr.*?</tr>`)
	allString := trReg.FindAllString(html, -1)
	tdReg := regexp.MustCompile(`<td.*?>(.*?)</td>`)
	for _, v := range allString {
		submatch := tdReg.FindAllStringSubmatch(v, -1)
		for _, v := range submatch {
			if len(v) == 2 {
				fmt.Println(v[1])
			}
		}
	}
}
