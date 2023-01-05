package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"github.com/goccy/go-json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"strconv"
	"time"
)

type Qinfo struct {
	Id    int    `gorm:"primary_key"`
	Qn    string `json:"qq"`
	Phone string `json:"phone"`
}

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:abc123..@tcp(127.0.0.1:3306)/qq?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 10)
}

func (Qinfo) TableName() string {
	return "qinfo"
}

func main() {
	var infos []Qinfo
	for count := 15; count < 19; count++ {
		file, err := ioutil.ReadFile("D:\\Q\\q\\" + strconv.FormatInt(int64(count), 10) + ".json")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(time.Now().UnixNano())
		err = json.Unmarshal(file, &infos)
		fmt.Println(time.Now().UnixNano())
		if err == nil {
			lenth := len(infos) / 500
			for i := 0; i < 500; i++ {
				if i != 499 {
					err = SaveBatch(infos[i*lenth : lenth*(i+1)])
				} else {
					err = SaveBatch(infos[i*lenth:])
				}
				if err != nil {
					fmt.Println("滑稽:", err)
				} else {
					fmt.Println("第", count, "--", i, "次存储成功")
				}
			}
		} else {
			fmt.Println("滑稽2", err)
		}
	}
}

func SaveBatch(infos []Qinfo) error {
	tx := db.Begin()
	var buffer bytes.Buffer
	sql := "insert into `qinfo`(`qn`,`phone`) values "
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for k, v := range infos {
		if k == len(infos)-1 {
			buffer.WriteString(fmt.Sprintf("('%s',%s);", v.Qn, v.Phone))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s'),", v.Qn, v.Phone))
		}
	}
	exec := db.Exec(buffer.String())
	affected := exec.RowsAffected
	err := exec.Error
	if err != nil || affected == 0 {
		fmt.Println(err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	exec.Row()
	return err
}
