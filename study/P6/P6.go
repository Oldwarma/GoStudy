package main

import (
	"bytes"
	"github.com/goccy/go-json"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	api := "http://127.0.0.1:8885/test"
	var data ApiData
	data.ApiType = 2
	data.RobotQQ = 1720538267
	data.GroupID = 154273490
	data.MessageType = 2
	data.MsgID = -1
	data.Message = "测试"
	data.UserID = 0
	requestPost(data, api)
}

type ApiData struct {
	ApiType     int64  `json:"api_type"`
	MsgID       int64  `json:"msg_id"`
	RobotQQ     int64  `json:"robot_qq"`
	Message     string `json:"message"`
	MessageType int64  `json:"message_type"`
	UserID      int64  `json:"user_id"`
	GroupID     int64  `json:"group_id"`
	MemberID    int64  `json:"member_id"`
	MessageNum  int64  `json:"message_num"`
	MessageID   int64  `json:"message_id"`
	Time        int64  `json:"time"`
	SubType     int64  `json:"sub_type"`
	RawMessage  string `json:"raw_message"`
	RejectMsg   string `json:"reject_msg"`
	Approve     int64  `json:"approve"`
	//Members map[string][]Member `json:"members"`
	GroupsID []string `json:"groups_id"`
}

func requestPost(data interface{}, url string) (bool, string) {
	bytesData, _ := json.Marshal(data)
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return false, "err"
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	if err != nil {
		return false, "err"
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, "err"
	}

	content := string(respBytes)
	return true, content
}
