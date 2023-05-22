package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Student struct {
	Name   string
	Age    int
	Gender string
	Score  float64
}

func main() {
	students := []Student{
		{"Tom", 18, "Male", 90.5},
		{"Mary", 19, "Female", 88.0},
		{"John", 20, "Male", 92.5},
	}
	f := excelize.NewFile()

	// 创建一个名为 "Sheet1" 的工作簿
	index := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	// 写入表头
	headers := []string{"Name", "Age", "Gender", "Score"}
	for col, header := range headers {
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'A'+col, 1), header)
	}

	// 写入数据
	for row, student := range students {
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'A', row+2), student.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'B', row+2), student.Age)
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'C', row+2), student.Gender)
		f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'D', row+2), student.Score)
	}
	// 保存文件
	err := f.SaveAs("students.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
