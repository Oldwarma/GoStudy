package timeUtil

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type LayoutTime time.Time

// 实现它的json序列化方法

func (this LayoutTime) MarshalJSON() ([]byte, error) {

	var stamp = fmt.Sprintf("\"%s\"", string(time.Time(this).Format(Layout))) // Format内即是你想转换的格式
	fmt.Println(stamp, "==")
	return []byte(stamp), nil

}

type LayoutTimeAAA struct {
	time.Time
}

func (t LayoutTimeAAA) MarshalJSON() ([]byte, error) {
	//格式化秒
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format(LayoutAAA))
	//seconds := t.Unix()
	//t.Format()
	return []byte(stamp), nil

}

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒

	seconds := t.Unix()

	return []byte(strconv.FormatInt(seconds, 10)), nil

}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time

	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil

	}

	return t.Time, nil

}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)

	if ok {
		*t = LocalTime{Time: value}

		return nil

	}

	return fmt.Errorf("can not convert %v to timestamp", v)

}

type TimeStamp int64

const CUS_TIME_FORMAT = "2006-01-02 15:04:05"

// 时间戳转换成日期字符串
func (ts TimeStamp) MarshalJSON() ([]byte, error) {
	t := time.Unix(int64(ts), 0)
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	b := make([]byte, 0, len(CUS_TIME_FORMAT)+2)
	//b = append(b, '"')
	b = t.AppendFormat(b, CUS_TIME_FORMAT)
	//b = append(b, '"')
	return b, nil
}

func (ts *TimeStamp) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	//parseTime, err := time.Parse(`"` + CUS_TIME_FORMAT + `"`, string(data))
	parseTime, err := time.Parse(CUS_TIME_FORMAT, string(data))
	if err != nil {
		return err
	}
	*ts = TimeStamp(parseTime.Unix())
	return nil
}
