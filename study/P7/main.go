package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{"reqName":"ContactList.AllGroupInfo","type":"resp","respCode":-99,"reqId":0,"resp":"卧槽"}`

	var a SnowAllMember
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	fmt.Println(a.Type)
	fmt.Println(a.ReqName)
	fmt.Println(a.Resp)
}

type SnowAllMember struct {
	ReqName  string          `json:"reqName"`
	Type     string          `json:"type"`
	RespCode int64           `json:"respCode"`
	ReqID    int64           `json:"reqId"`
	Resp     map[string][]ST `json:"resp"`
}

type ST struct {
	ID      int `json:"id"`
	IsAdmin int `json:"isAdmin"`
}
