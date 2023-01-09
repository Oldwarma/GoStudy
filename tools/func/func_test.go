package _func

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

type A struct {
	Name string
	Age  string
}

type B struct {
	Name1 string
	Age   string
}

func TestAAA(t *testing.T) {
	fmt.Println(SplitArray([]string{"12", "22", "333"}, 2))
}

func TestGetDeviceId(t *testing.T) {
	newUUID, err := uuid.NewUUID()
	fmt.Println(newUUID)
	fmt.Println(err)
}
