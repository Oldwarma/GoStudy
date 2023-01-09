package response

import (
	"gostudy/errno"
	"reflect"
)

// 数据返回通用数据结构
type RpcResponse struct {
	Code    string      `json:"code"`    // 错误码
	Message string      `json:"message"` // 提示信息
	Success bool        `json:"success"` // ((0:成功, 1:失败, >1:错误码))
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
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

func DeleteError(in interface{}, message string) (out interface{}) {
	v := reflect.ValueOf(in).Elem()
	tmp := reflect.New(v.Type()).Elem()
	tmp.Set(v)
	tmp.FieldByName("Code").SetString(errno.FailureDbDelete.ErrorCode)
	tmp.FieldByName("Success").SetBool(false)
	tmp.FieldByName("Message").SetString(errno.FailureDbDelete.Note + "_" + message)
	v.Set(tmp)
	return in
}

func FindError(in interface{}, message string) (out interface{}) {
	v := reflect.ValueOf(in).Elem()
	tmp := reflect.New(v.Type()).Elem()
	tmp.Set(v)
	tmp.FieldByName("Code").SetString(errno.FailureDbFind.ErrorCode)
	tmp.FieldByName("Success").SetBool(false)
	tmp.FieldByName("Message").SetString(errno.FailureDbFind.Note + "_" + message)
	v.Set(tmp)
	return in
}

// RcpToMessageInfo resDataDevice.Status = _func.BoolToString(i2.Success)
//					resDataDevice.Note = i2.Message
//					resDataDevice.ErrorCode = i2.Code
func RcpToMessageInfo(InRpc interface{}, InApi interface{}) {
	rpcOf := reflect.ValueOf(InRpc).Elem()
	success := rpcOf.FieldByName("Success").String()
	message := rpcOf.FieldByName("Message").String()
	code := rpcOf.FieldByName("Code").String()

	v := reflect.ValueOf(InApi).Elem()
	tmp := reflect.New(v.Type()).Elem()
	tmp.Set(v)
	tmp.FieldByName("Status").SetString(success)
	tmp.FieldByName("Note").SetString(message)
	tmp.FieldByName("ErrorCode").SetString(code)
	v.Set(tmp)
}
