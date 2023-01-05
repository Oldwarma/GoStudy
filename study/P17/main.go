package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

func main() {
	xlFile, err := xlsx.OpenFile("1.xlsx")
	if err != nil {
		fmt.Println("打开失败")
	}
	//中间数据，忽略
	investmentCompanyMap := make(map[string]bool)

	//投资的公司索引
	investmentCompanyIndex := make(map[string]int64)

	//投资的公司
	var investmentCompanys []string
	//创业企业  ------ 投资的公司
	investmentMap := make(map[string][]string)

	sheet := xlFile.Sheets[0]
	for k, row := range sheet.Rows {
		if k == 0 {
			continue
		}
		entrepreneurshipCompany := row.Cells[1].String()
		investmentCompany := row.Cells[2].String()
		if entrepreneurshipCompany == "" {
			continue
		}
		investmentCompanyMap[investmentCompany] = true
		investmentMap[entrepreneurshipCompany] = append(investmentMap[entrepreneurshipCompany], investmentCompany)
	}
	//排序
	for k, _ := range investmentCompanyMap {
		investmentCompanys = append(investmentCompanys, k)
	}
	var row *xlsx.Row
	var cell *xlsx.Cell

	sheet, err = xlFile.AddSheet("处理数据")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(0.6)
	for k, v := range investmentCompanys {
		investmentCompanyIndex[v] = int64(k)
		cell = row.AddCell()
		if k == 0 {
			cell = row.AddCell()
		}
		cell.Value = v
	}
	for _, v := range investmentCompanys {
		row = sheet.AddRow()
		row.SetHeightCM(0.6)
		cell = row.AddCell()
		cell.Value = v
	}
	//楼上构造表格基础

	//A 和 B 投资的 共同数量
	data := make(map[string]map[string]int64)

	//投资的公司
	for _, v := range investmentCompanys {
		for _, v2 := range investmentCompanys {
			for _, v3 := range investmentMap {
				if ContainPro(v3, v, v2) {
					if data[v] == nil {
						data[v] = make(map[string]int64)
					}
					data[v][v2] = data[v][v2] + 1
				}
			}
		}
	}

	okTemp := make(map[string]map[string]bool)

	for k, r := range sheet.Rows {
		if k == 0 {
			continue
		}
		rowName := sheet.Rows[k].Cells[0].String()
		for k2, _ := range sheet.Rows[0].Cells {
			content := sheet.Rows[0].Cells[k2].String()
			if content == "" {
				continue
			}
			cell = r.AddCell()
			count := data[rowName][content]
			if rowName == content {
				count = 0
			}
			//else {
			//	count = 1
			//}
			cellOk := okTemp[content]
			if cellOk == nil {
				cellOk = make(map[string]bool)
			}

			if cellOk[rowName] {
				cell.Value = ""
			} else {
				if count == 0 {
					cell.Value = ""
				} else {
					count = 1
					cell.Value = strconv.FormatInt(count, 10)
				}
				if okTemp[rowName] == nil {
					okTemp[rowName] = make(map[string]bool)
				}
				okTemp[rowName][content] = true
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

func ContainPro(arr []string, data ...string) bool {
	for _, v := range data {
		if !Contain(v, arr) {
			return false
		}
	}
	return true
}
