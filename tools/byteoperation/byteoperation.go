package byteoperation

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"
	"unsafe"
)

//BytesToIntB byte转int  isSymbol是否带符号
func BytesToIntB(b []byte, isSymbol bool) int {
	if isSymbol {
		i, err := bytesToIntS(b, binary.BigEndian)
		if err != nil {
			fmt.Printf("BytesToInt err %v", err)
		}
		return i
	}
	i, err := bytesToIntU(b, binary.BigEndian)
	if err != nil {
		fmt.Printf("BytesToInt err %v", err)
	}
	return i
}

func BytesToIntL(b []byte, isSymbol bool) int {
	if isSymbol {
		i, err := bytesToIntS(b, binary.LittleEndian)
		if err != nil {
			fmt.Printf("BytesToInt err %v", err)
		}
		return i
	}
	i, err := bytesToIntU(b, binary.LittleEndian)
	if err != nil {
		fmt.Printf("BytesToInt err %v", err)
	}
	return i
}

//字节数组转成int(无符号的)
func bytesToIntU(b []byte, order binary.ByteOrder) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, order, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, order, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, order, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

//字节数组转成int(有符号)
func bytesToIntS(b []byte, order binary.ByteOrder) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, order, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, order, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, order, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

// IntToBytes int转byte切片
func IntToBytes(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

func IntToBytesL(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

// StrToBytes 字符串转byte切片
func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToStr byte切片转字符串
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytesCustomLen 根据字符串生成自定义字节的byte切片，不够补0
func StringToBytesCustomLen(gatewayId string, byteLength int) ([]byte, error) {
	gaByte := []byte(gatewayId)
	if len(gaByte) > byteLength {
		return nil, errors.New("gateway id  through the long !!!")
	} else {
		var tempByte []byte
		for i := 0; i < byteLength-len(gaByte); i++ {
			tempByte = append(tempByte, 0x00)
		}
		tempByte = append(gaByte, tempByte...)
		return tempByte, nil
	}
}

//uint16转数组
func Uint16ToBytes(value ...uint16) []byte {
	data := make([]byte, 2*len(value))
	for i, v := range value {
		binary.BigEndian.PutUint16(data[i*2:], v)
	}
	return data
}

func RandBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	b := make([]byte, length)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return nil
	}
	return b
}

func GatewayId(gatewayId string) []byte {
	return append(StrToBytes(gatewayId), 0x00)
}

func RandMid() string {
	result, _ := rand.Int(rand.Reader, big.NewInt(4294967295))
	return result.String()
}
