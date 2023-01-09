package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Md5WithDir(dir string) string {
	bytes := ReadFile(dir)
	if len(bytes) == 0 {
		return ""
	}
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
