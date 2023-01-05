// 代码由简易GoVCL IDE自动生成。
// 不要更改此文件名
// 在这里写你的事件。

package main

import (
	"github.com/ying32/govcl/vcl"
	"strconv"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	for i := 0; i < 20; i++ {
		f.FirstWeekStart.Items().Add(strconv.FormatInt(int64(i), 10))
		f.FirstWeekEnd.Items().Add(strconv.FormatInt(int64(i), 10))
		f.SecondWeekStart.Items().Add(strconv.FormatInt(int64(i), 10))
		f.SecondWeekEnd.Items().Add(strconv.FormatInt(int64(i), 10))
		f.ThirdWeekStart.Items().Add(strconv.FormatInt(int64(i), 10))
		f.ThirdWeekEnd.Items().Add(strconv.FormatInt(int64(i), 10))
	}
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {

	var firstStart int64 = 0
	if f.FirstWeekStart.Items().Text() != "" {
		firstStart, _ = strconv.ParseInt(f.FirstWeekStart.Items().Text(), 10, 64)
	}

	var firstEnd int64 = 0
	if f.FirstWeekEnd.Items().Text() != "" {
		firstEnd, _ = strconv.ParseInt(f.FirstWeekEnd.Items().Text(), 10, 64)
	}

	var secondStart int64 = 0
	if f.SecondWeekStart.Items().Text() != "" {
		secondStart, _ = strconv.ParseInt(f.SecondWeekStart.Items().Text(), 10, 64)
	}

	var secondEnd int64 = 0
	if f.SecondWeekEnd.Items().Text() != "" {
		secondEnd, _ = strconv.ParseInt(f.SecondWeekEnd.Items().Text(), 10, 64)
	}

	var thirdStart int64 = 0
	if f.ThirdWeekStart.Items().Text() != "" {
		thirdStart, _ = strconv.ParseInt(f.ThirdWeekStart.Items().Text(), 10, 64)
	}

	var thirdEnd int64 = 0
	if f.ThirdWeekEnd.Items().Text() != "" {
		thirdEnd, _ = strconv.ParseInt(f.ThirdWeekEnd.Items().Text(), 10, 64)
	}

	min := MinInt(firstStart, secondStart, thirdStart)
	max := MaxInt(firstEnd, secondEnd, thirdEnd)

	f.FirstWeekStart.Items().Clear()
	f.FirstWeekEnd.Items().Clear()
	f.SecondWeekStart.Items().Clear()
	f.SecondWeekEnd.Items().Clear()
	f.ThirdWeekStart.Items().Clear()
	f.ThirdWeekEnd.Items().Clear()

	var freeFitstStart []int64
	var freeFitsrEnd []int64

	var freeSecondStary []int64
	var freeSecondEnd []int64

	var freeThirdStart []int64
	var freeThirdEnd []int64

}

func MaxInt(data ...int64) int64 {
	var max int64 = 0
	for k, v := range data {
		if k == 0 {
			max = v
		} else {
			if v > max {
				max = v
			}
		}
	}
	return max
}

func MinInt(data ...int64) int64 {
	var min int64 = 0
	for k, v := range data {
		if k == 0 {
			min = v
		} else {
			if v < min {
				min = v
			}
		}
	}
	return min
}
