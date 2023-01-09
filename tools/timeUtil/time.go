package timeUtil

import (
	"strconv"
	"time"
)

const (
	Layout         = "2006-01-02 15:04:05"
	LayoutAAA      = "2006-01-02 15:04:05.000"
	LayoutNoSecond = "2006-01-02 15:04"
)

func GetNowTimeFormat() string {
	return time.Now().Format(Layout)
}

func GetNowTimeFormatAAA() string {
	return time.Now().Format(LayoutAAA)
}

func GetTimeFormatAAA(t time.Time) string {
	return t.Format(LayoutAAA)
}

func GetTimeFormat(t time.Time) string {
	return t.Format(Layout)
}

func GetTimeStringFormat(str string) string {
	loc, _ := time.LoadLocation("Local")
	tt, _ := time.ParseInLocation(Layout, str, loc)
	return tt.String()
}

func GetTimeParseAAA(t string) int64 {
	loc, _ := time.LoadLocation("Local")
	tt, err := time.ParseInLocation(LayoutAAA, t, loc)
	if err != nil {
		return 0
	}
	return tt.Unix()
}

// 获取秒
func GetScTimeFormatAAA(t time.Time) string {
	return t.Format(Layout)
}

func GetNowScTime() int64 {
	return time.Now().Unix()
}

func GetNowScTimeString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func GetScTimeParseAAA(t string) int64 {
	loc, _ := time.LoadLocation("Local")
	tt, err := time.ParseInLocation(Layout, t, loc)
	if err != nil {
		return 0
	}
	return tt.Unix()
}
