package crc

import (
	"hash/crc32"
	"io/ioutil"
	"os"
	"strconv"
)

// CRC32 文件检验码，算法采用CRC32，格式采用16进制，如“0x12345678”
func CRC32(path string) (string, error) {
	fileobj, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fileobj.Close()
	all, err := ioutil.ReadAll(fileobj)
	if err != nil {
		return "", err
	}
	checksumIEEE := crc32.ChecksumIEEE(all)
	return "0x" + strconv.FormatUint(uint64(checksumIEEE), 16), nil
}
