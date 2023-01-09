package _func

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// null
func NulRetStr(str interface{}) interface{} {
	if str == nil {
		return ""
	}
	return str
}

func GetDbId() string {
	id := uuid.NewString()
	return strings.ReplaceAll(id, "-", "")
}

// 设备id
func DevUUid() string {
	id, _ := uuid.New().Time().UnixTime()
	return gconv.String(id)
}

// 模型id
func ProUUid() string {
	id, _ := uuid.New().Time().UnixTime()
	return gconv.String(id)
}

// 模型id
func ModUUid() string {
	id, _ := uuid.New().Time().UnixTime()
	return gconv.String(id)
}

// 服务id
func SerUUid() string {
	id, _ := uuid.New().Time().UnixTime()
	return gconv.String(id)
}

func IntToStr(str int) string {
	var i = strconv.Itoa(str)
	return i
}
func StrToInt(str string) int {
	var i, _ = strconv.Atoi(str)
	return i
}

// 属性id
func PropUUid() string {
	id, _ := uuid.New().Time().UnixTime()
	return gconv.String(id)
}

// 命令id
func CommUUid() string {
	id, _ := uuid.New().Time().UnixTime()
	return gconv.String(id)
}

// 命令属性id
func commPropUUid() string {
	id, _ := uuid.New().Time().UnixTime()
	return gconv.String(id)
}

// 数据id 与数据 name 对应
func IdToName(id *string, name *string, types string) {
	name = id
	switch types {
	case "dev":
		rid := DevUUid()
		id = &rid

	case "pro":
		rid := ProUUid()
		id = &rid

	case "mod":
		rid := ModUUid()
		id = &rid

	case "prop":
		rid := PropUUid()
		id = &rid

	case "comm":
		rid := CommUUid()
		id = &rid

	case "comm_pro":
		rid := commPropUUid()
		id = &rid

	default:
		rid := DevUUid()
		id = &rid

	}
	//return name, id
}

// name转id
func NameToId(id, name *string) {
	name = id
}

// 是否数组
func IsArray(key string, arr []string) bool {
	for _, s := range arr {
		if s == key {
			return true
		}
	}
	return false
}

// 是否相等
func IsEqual(key string, arr []string) bool {
	for _, s := range arr {
		if s == key {
			return true
		}
	}
	return false
}

// 是否相等
func IsMapEqual(key string, arr []string) bool {
	for _, s := range arr {
		if s == key {
			return true
		}
	}
	return false
}

// nil 与 空
func IsEmpty(key interface{}) bool {
	if key == nil || key == "" || key == 0 {
		return true
	}
	return false
}

// 结构体转map
func StructToMap(str interface{}) (map[string]interface{}, map[string]interface{}) {
	ob := reflect.TypeOf(str)
	obV := reflect.ValueOf(str)
	var res = make(map[string]interface{}, 0)
	var resLower = make(map[string]interface{}, 0)
	for i := 0; i < ob.NumField(); i++ {
		fieldType := ob.Field(i).Name
		fieldValue := obV.Field(i)
		if !IsEmpty(fieldValue) {
			res[fieldType] = fieldValue
			resLower[StringToLower(fieldType)] = fieldValue
		}
	}
	return res, res
}

// 驼峰转小写下划线
func StringToLower(name string) string {
	buffer := strings.Builder{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteString("_")
			}
			buffer.WriteString(gconv.String(unicode.ToLower(r)))
		} else {
			buffer.WriteString(gconv.String(r))
		}
	}
	return buffer.String()
}

// 结构体与map值自动赋值给新结构体
func StructMapToNewStruct(new interface{}) {

}

// 结构体转数组
func StructToStringArray(str interface{}, tag string) []string {
	structType := reflect.TypeOf(str).Kind()
	structVal := reflect.ValueOf(str)
	var strArr []string
	if structType == reflect.Slice {
		for i := 0; i < structVal.Len(); i++ {
			strArr = append(strArr, structVal.Index(i).FieldByName(tag).String())
		}
	}
	return strArr
}

func IfStrElseStr(IfKey string, IfValue string, value string) string {
	if IfKey == IfValue {
		return value
	}
	return ""
}

// 结构体转json
func StructToJson(str interface{}) string {
	b, _ := json.Marshal(str)
	return string(b)
}

// json转结构体
func JsonToStruct(js interface{}, res interface{}) {
	bt := gconv.Bytes(js)
	_ = json.Unmarshal(bt, &res)
}

// 结构体拼接赋值
func GoStruct(params interface{}, pointer interface{}, keyKey map[string]string) {
	paramsVal := reflect.ValueOf(params)
	pointerVal := reflect.ValueOf(pointer)
	fmt.Println(paramsVal.Len())
	fmt.Println(pointerVal.Len())
	for i := 0; i < paramsVal.Len(); i++ {
		fmt.Println(paramsVal.Index(i).Elem())
	}
}

func BoolToString(b bool) string {
	if b == true {
		return "0"
	}
	return "1"
}

// rpc 返回
func AssembleMessage(funName string, data interface{}) interface{} {
	message := struct {
		Head struct {
			FunName string `json:"FunName,omitempty"`
		} `json:"Head"`
		Data interface{} `json:"Data,omitempty"`
	}{}
	message.Head.FunName = funName
	message.Data = data
	return &message
}

// 数组平分
func SplitArray(arr []string, num int64) [][]string {
	var segments = make([][]string, 0)
	l := len(arr)
	i := int64(0)
	for true {
		if int64(l)-i < num {
			segments = append(segments, arr[i:])
			break
		}
		segments = append(segments, arr[i:i+num])
		i += num
	}
	return segments
}

// InterfaceToString Interface 转 string
func InterfaceTo(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func StringToInterface(in []string) (res []interface{}) {
	for _, s := range in {
		res = append(res, s)
	}
	return res
}
