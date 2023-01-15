package main

import (
	"fmt"
	"strings"
)

const st = "你干嘛哎呦 荔枝一点 hello CXK"

var arr []string

func main() {
	arr = append(arr, "ToUpper:"+strings.ToUpper(st))
	arr = append(arr, "ToLower:"+strings.ToLower(st))
	arr = append(arr, "ToTitle:"+strings.ToTitle(st))
	arr = append(arr, "Repeat:"+strings.Repeat(st, 5))
	arr = append(arr, "TrimSpace:"+strings.TrimSpace(st)) //头尾
	arr = append(arr, "TrimLeft:"+strings.TrimLeft(st, "点"))
	arr = append(arr, "Replace:"+strings.Replace(st, "你干嘛哎呦", "niganma", 2))

	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("Split:", strings.Split(st, " "))
	fmt.Println("TrimSpace:", strings.Compare(st, st))
	fmt.Println("Count:", strings.Count(st, "l"))
	fmt.Println("HasPrefix:", strings.HasPrefix(st, st)) //头
	fmt.Println("HasSuffix:", strings.HasSuffix(st, st)) //尾
	fmt.Println("EqualFold:", strings.EqualFold(st, st))
}
