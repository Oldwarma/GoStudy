package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	output, err := xlsx.FileToSlice("./f.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var cName []string
	rows := output[0][0]
	for _, row := range rows {
		cName = append(cName, row)
	}

	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("data")
	row := sheet.AddRow()

	//初始化标题
	row = sheet.AddRow()
	for _, v := range cName {
		cell := row.AddCell()
		cell.Value = v
	}

	//开始造数据
	for i := 0; i < 10000; i++ {
		row = sheet.AddRow()
		var sum int64
		for k, _ := range cName {
			if k == 0 {
				cell := row.AddCell()
				cell.Value = RandX(1, 2, 50, 100)
			}
			if k == 1 || k == 3 || k == 4 {
				cell := row.AddCell()
				cell.Value = RandX(1, 4, 10, 50, 80, 100)
			}
			if k == 2 {
				cell := row.AddCell()
				cell.Value = RandX(1, 5, 0, 10, 50, 80, 100)
			}
			if k == 5 {
				cell := row.AddCell()
				cell.Value = RandX(1, 2, 40, 100)
			}
			if k >= 6 && k <= 11 {
				cell := row.AddCell()
				cell.Value = RandX(0, 1, 30, 100)
			}
			if k >= 12 && k <= 16 {
				cell := row.AddCell()
				cell.Value = RandX(0, 1, 20, 100)
			}
			if k >= 17 && k <= 24 {
				cell := row.AddCell()
				cell.Value = RandX(0, 1, 50, 100)
			}
			if k >= 25 && k <= 28 {
				cell := row.AddCell()
				cell.Value = RandX(0, 1, 50, 100)
			}
			if k >= 29 && k <= 33 {
				cell := row.AddCell()
				cell.Value = RandX(0, 1, 40, 100)
			}
			if k >= 34 && k <= 38 {
				cell := row.AddCell()
				cell.Value = RandX(0, 1, 40, 100)
			}
			if k >= 39 && k <= 49 {
				cell := row.AddCell()
				x := RandX(1, 5, 10, 30, 50, 80, 100)
				cell.Value = x
				parseInt, _ := strconv.ParseInt(x, 10, 64)
				sum = sum + parseInt
			}
		}
		cell := row.AddCell()
		cell.Value = strconv.FormatInt(sum, 10)
	}
	err = file.Save("./data.xlsx")
	fmt.Println(err)
}

func RandX(min, max int64, rate ...int64) string {
	//num := rand.Int63n(max-min) + min
	n := rand.Int63n(100)
	for k, v := range rate {
		if n < v {
			return strconv.FormatInt(min+int64(k), 10)
		}
	}
	return strconv.FormatInt(max, 10)
}
