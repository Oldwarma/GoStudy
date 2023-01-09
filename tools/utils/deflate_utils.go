package utils

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"errors"
	"io"
	"io/ioutil"
)

func UnDeflate(data []byte) []byte {
	reader := bytes.NewReader(data)
	flateReader := flate.NewReader(reader)
	defer flateReader.Close()
	//copy flate Reader中的内容
	deBuffer := new(bytes.Buffer)
	_, err := io.Copy(deBuffer, flateReader)
	if err != nil {
		return nil
	}
	return deBuffer.Bytes()
}

func UnZlib(data []byte) []byte {
	buf := bytes.NewReader(data)
	reader, err := zlib.NewReader(buf)
	if err != nil {
		return nil
	}
	defer reader.Close()
	content, err := ioutil.ReadAll(reader)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
		return nil
	}
	return content
}
