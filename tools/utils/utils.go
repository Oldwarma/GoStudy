package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

func Bytes2MbStr(total uint64) string {
	return fmt.Sprintf("%d", total/1024/1024)
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func UInt162Bytes(seq uint16) []byte {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, seq)
	return bs
}

func Int28Bytes(num int64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(num))
	return bs
}

func Bytes2IntWithLit(bys []byte) int {
	switch len(bys) {
	case 4:
		bytebuffer := bytes.NewBuffer(bys)
		var data uint32
		binary.Read(bytebuffer, binary.LittleEndian, &data)
		return int(data)
	case 2:
		bytebuffer := bytes.NewBuffer(bys)
		var data uint16
		binary.Read(bytebuffer, binary.LittleEndian, &data)
		return int(data)
	case 1:
		bytebuffer := bytes.NewBuffer(bys)
		var data uint8
		binary.Read(bytebuffer, binary.LittleEndian, &data)
		return int(data)
	}
	return 0
}

func Bytes2Int(bys []byte) int {
	switch len(bys) {
	case 4:
		bytebuffer := bytes.NewBuffer(bys)
		var data uint32
		binary.Read(bytebuffer, binary.BigEndian, &data)
		return int(data)
	case 2:
		bytebuffer := bytes.NewBuffer(bys)
		var data uint16
		binary.Read(bytebuffer, binary.BigEndian, &data)
		return int(data)
	case 1:
		bytebuffer := bytes.NewBuffer(bys)
		var data uint8
		binary.Read(bytebuffer, binary.BigEndian, &data)
		return int(data)
	}
	return 0
}

func Int22Bytes(n int) []byte {
	data := uint16(n)
	bytebuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytebuffer, binary.BigEndian, data)
	return bytebuffer.Bytes()
}

func IntToBytes(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		temp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &temp)
		return bytesBuffer.Bytes(), nil
	case 2:
		temp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &temp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		temp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &temp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

func BytesToIntB(b []byte, isSymbol bool) int {
	if isSymbol {
		i, err := BytesToIntS(b, binary.BigEndian)
		if err != nil {
			fmt.Errorf("BytesToInt err %v", err)
		}
		return i
	}
	i, err := BytesToIntU(b, binary.BigEndian)
	if err != nil {
		fmt.Errorf("BytesToInt err %v", err)
	}
	return i
}

func BytesToIntS(b []byte, order binary.ByteOrder) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var temp int8
		err := binary.Read(bytesBuffer, order, &temp)
		return int(temp), err
	case 2:
		var temp int16
		err := binary.Read(bytesBuffer, order, &temp)
		return int(temp), err
	case 4:
		var temp int32
		err := binary.Read(bytesBuffer, order, &temp)
		return int(temp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")

	}
}

func BytesToIntU(b []byte, order binary.ByteOrder) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var temp uint8
		err := binary.Read(bytesBuffer, order, &temp)
		return int(temp), err
	case 2:
		var temp uint16
		err := binary.Read(bytesBuffer, order, &temp)
		return int(temp), err
	case 4:
		var temp uint32
		err := binary.Read(bytesBuffer, order, &temp)
		return int(temp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}

}

func UInt642Bytes(seq uint64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, seq)
	return bs
}

func Int2BytesWithLit(i int, s uint8) []byte {
	b := make([]byte, s)
	for k, _ := range b {
		b[k] = byte(i >> (k * 8) & 0xff)
	}
	return b
}
