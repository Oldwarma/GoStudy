package helper

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func Debug(tag string, any ...interface{}) {
	fmt.Printf("_%s\n", tag)
	g.Dump(any)
	fmt.Printf("%s_\n", tag)
}

func IIF(condiction bool, x, y interface{}) interface{} {
	if condiction {
		return x
	}

	return y
}
