package main

import (
	"fmt"
	"gostudy/errno"
	"reflect"
)

type resp struct {
	Code    string
	Success bool
	Message string
}

func main() {
	var req resp
	Success(&req, "")
	fmt.Println(req)
}

// Success 返回成功
func Success(in interface{}, data interface{}) (out interface{}) {
	v := reflect.ValueOf(in).Elem()
	tmp := reflect.New(v.Type()).Elem()
	tmp.Set(v)
	tmp.FieldByName("Code").SetString(errno.Success.ErrorCode)
	tmp.FieldByName("Success").SetBool(true)
	tmp.FieldByName("Message").SetString(errno.Success.Note)
	//tmp.FieldByName("Data").SetString(data)
	v.Set(tmp)
	return in
}

// Error 返回错误结构
func CreateError(in interface{}, message string) (out interface{}) {
	v := reflect.ValueOf(in).Elem()
	tmp := reflect.New(v.Type()).Elem()
	tmp.Set(v)
	tmp.FieldByName("Code").SetString(errno.FailureDbInsert.ErrorCode)
	tmp.FieldByName("Success").SetBool(false)
	tmp.FieldByName("Message").SetString(errno.FailureDbInsert.Note + "_" + message)
	v.Set(tmp)
	return in
}
