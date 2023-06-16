package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := int64(1622266928) // 输入你的时间戳
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05")) // 修改格式化字符串以适应需要的日期时间格式
}
