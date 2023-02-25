package main

import "unsafe"

func main() {

}

func ByteSlice2String(b []byte) string {
	return unsafe.String(&b[0], len(b))
}
func String2BytesSlice(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
