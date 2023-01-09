package utils

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func Bcd2Number(bcd []byte) string {
	var number string
	/*	for _, i := range bcd {
		number += fmt.Sprintf("%02X", i)
	}*/
	for i := len(bcd) - 1; i >= 0; i-- {
		number += fmt.Sprintf("%02X", bcd[i])
	}

	pos := strings.LastIndex(number, "F")
	if pos == 8 {
		return "0"
	}
	return number[pos+1:]
}

func byte2ToFloat3(bcd []byte) string {
	bits := float64(int16(binary.BigEndian.Uint16(bcd)))
	return fmt.Sprintf("%.3f", bits/1000.0)
}

func byte2ToFloat2(bcd []byte) string {
	bits := float64(binary.BigEndian.Uint16(bcd))
	return fmt.Sprintf("%.2f", bits/100.0)
}

func Bcd2ToFloat1(bcd []byte) string {
	data := make([]byte, 0)
	data = append(data, bcd[1]>>4)
	data = append(data, bcd[1]&0x0f)
	data = append(data, bcd[0]>>4)
	data = append(data, bcd[0]&0x0f)
	return fmt.Sprintf("%d%d%d.%d", data[0], data[1], data[2], data[3])
}

func Bcd2ToFloat2(bcd []byte) string {
	data := make([]byte, 0)
	data = append(data, bcd[1]>>4)
	data = append(data, bcd[1]&0x0f)
	data = append(data, bcd[0]>>4)
	data = append(data, bcd[0]&0x0f)
	return fmt.Sprintf("%d%d.%d%d", data[0], data[1], data[2], data[3])
}

func Bcd3ToFloat3(bcd []byte) string {
	data := make([]byte, 0)
	data = append(data, bcd[2]>>4)
	data = append(data, bcd[2]&0x0f)
	data = append(data, bcd[1]>>4)
	data = append(data, bcd[1]&0x0f)
	data = append(data, bcd[0]>>4)
	data = append(data, bcd[0]&0x0f)
	return fmt.Sprintf("%d%d%d.%d%d%d", data[0], data[1], data[2],
		data[3], data[4], data[5])
}

func Bcd4ToFloat2(bcd []byte) string {
	data := make([]byte, 0)
	data = append(data, bcd[3]>>4)
	data = append(data, bcd[3]&0x0f)
	data = append(data, bcd[2]>>4)
	data = append(data, bcd[2]&0x0f)
	data = append(data, bcd[1]>>4)
	data = append(data, bcd[1]&0x0f)
	data = append(data, bcd[0]>>4)
	data = append(data, bcd[0]&0x0f)
	return fmt.Sprintf("%d%d%d%d%d%d.%d%d", data[0], data[1], data[2],
		data[3], data[4], data[5], data[6], data[7])
}

func Bcd2ToFloat4(bcd []byte) string {
	data := make([]byte, 0)
	data = append(data, bcd[2]>>4)
	data = append(data, bcd[2]&0x0f)
	data = append(data, bcd[1]>>4)
	data = append(data, bcd[1]&0x0f)
	data = append(data, bcd[0]>>4)
	data = append(data, bcd[0]&0x0f)
	return fmt.Sprintf("%d%d.%d%d%d%d", data[0], data[1], data[2],
		data[3], data[4], data[5])
}

func BcdToBin(bcd byte) byte {
	return (bcd>>4)*10 + (bcd & 0x0f)
}
