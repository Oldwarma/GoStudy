package timeUtil

import (
	"testing"
	"time"
)

func TestGetNowTimeFormat(t *testing.T) {
	t.Log(GetNowTimeFormat())
	t.Log(GetNowTimeFormatAAA())
	t.Log(GetTimeFormatAAA(time.Now()))
}

func TestGetTimeFormatAAA(t *testing.T) {
	t.Log(GetTimeFormatAAA(time.Now()))
}

func TestGetTimeParseAAA(t *testing.T) {
	t.Log(GetTimeParseAAA("2022-03-29 11:37:53.292"))
	t.Log(GetTimeParseAAA("2022-03-28 20:06:38"))
	t.Log(GetTimeParseAAA("2022.03.28 20:06:38"))
	t.Log(GetTimeStringFormat("2022-07-22 16:02:27 +0800 CST"))
	t.Log(GetScTimeParseAAA("2022-07-22 16:02:27"))
	t.Log(GetScTimeParseAAA("2022-07-22 16:02:66"))
}
