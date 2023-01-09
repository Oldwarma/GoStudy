package errno

import (
	"bytes"
	"errors"
)

type Errno struct {
	ErrorCode string
	Note      string
}

func (err *Errno) Error() string {
	var showErrorText bytes.Buffer
	showErrorText.WriteString(err.ErrorCode)
	showErrorText.WriteString(":")
	showErrorText.WriteString(err.Note)
	return showErrorText.String()
}

func IsSuccess(errno *Errno) bool {
	return errors.Is(errno, Success)
}
