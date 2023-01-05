package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

func main() {
	xlFile, err := xlsx.OpenFile("1.xlsx")
	if err != nil {
		fmt.Println("打开失败")
	}

	//领投  ---- 跟投
	investMap := make(map[string][]string)

	//创业公司 --- 领投--- 跟投
	EntrepreneurshipMap := make(map[string]map[string][]string)

	//全部公司
	allCompany := make(map[string]bool)

	sheet := xlFile.Sheets[0]

	//初始化领投数据
	for k, row := range sheet.Rows {
		if k == 0 {
			continue
		}
		if len(row.Cells) < 4 || row.Cells == nil {
			continue
		}
		if row.Cells[0] == nil || row.Cells[1] == nil || row.Cells[2] == nil || row.Cells[3] == nil {
			continue
		}
		investType := row.Cells[3].String()
		investCompanyName := row.Cells[2].String()
		EntrepreneurshipName := row.Cells[1].String()
		allCompany[investCompanyName] = true
		iMap := EntrepreneurshipMap[EntrepreneurshipName]
		if iMap == nil {
			iMap = make(map[string][]string)
		}
		if strings.Contains(investType, "领") {
			var follows []string
			follows = iMap[investCompanyName]
			iMap[investCompanyName] = follows
		}
		EntrepreneurshipMap[EntrepreneurshipName] = iMap
	}

	//初始化跟投数据
	for k, row := range sheet.Rows {
		if k == 0 {
			continue
		}
		if len(row.Cells) < 4 || row.Cells == nil {
			continue
		}
		if row.Cells[0] == nil || row.Cells[1] == nil || row.Cells[2] == nil || row.Cells[3] == nil {
			continue
		}

		investType := row.Cells[3].String()
		investCompanyName := row.Cells[2].String()
		EntrepreneurshipName := row.Cells[1].String()

		if strings.Contains(investType, "跟") {
			iMap := EntrepreneurshipMap[EntrepreneurshipName]
			firstName := ""
			for name, _ := range iMap {
				firstName = name
				break
			}
			var follows []string
			follows = iMap[firstName]
			follows = append(follows, investCompanyName)
			iMap[firstName] = follows
			EntrepreneurshipMap[EntrepreneurshipName] = iMap
		}
	}

	for _, v := range EntrepreneurshipMap {
		for k2, v2 := range v {
			investMap[k2] = v2
		}
	}

	var row *xlsx.Row
	var cell *xlsx.Cell

	sheet, err = xlFile.AddSheet("处理数据")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(0.6)

	isFirst := true

	var cs []string
	for k, _ := range allCompany {
		if k == "" {
			continue
		}
		cs = append(cs, k)
	}
	for _, v := range cs {
		cell = row.AddCell()
		if isFirst {
			cell = row.AddCell()
			isFirst = false
		}
		cell.Value = v
	}

	for _, v := range cs {
		row = sheet.AddRow()
		row.SetHeightCM(0.6)
		cell = row.AddCell()
		cell.Value = v
	}
	//for k, _ := range allCompany {
	//	cell = row.AddCell()
	//	if isFirst {
	//		cell = row.AddCell()
	//		isFirst = false
	//	}
	//	cell.Value = k
	//}

	//for k, _ := range allCompany {
	//	row = sheet.AddRow()
	//	row.SetHeightCM(0.6)
	//	cell = row.AddCell()
	//	cell.Value = k
	//}
	//楼上构造表格基础

	for k, r := range sheet.Rows {
		if k == 0 {
			continue
		}
		firstName := sheet.Rows[k].Cells[0].String()
		follows := investMap[firstName]
		for k2, v2 := range sheet.Rows[0].Cells {
			if k2 == 0 || v2.String() == "" {
				continue
			}
			cell = r.AddCell()
			if Contain(v2.String(), follows) {
				cell.Value = "1"
			} else {
				cell.Value = ""
			}
		}
	}
	err = xlFile.Save("1.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func Contain(data string, arr []string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v == data {
			return true
		}
	}
	return false
}
