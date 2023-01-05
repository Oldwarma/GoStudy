package main

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize              ulong
	cntUsage            ulong
	th32ProcessID       ulong
	th32DefaultHeapID   ulong_ptr
	th32ModuleID        ulong
	cntThreads          ulong
	th32ParentProcessID ulong
	pcPriClassBase      ulong
	dwFlags             ulong
	szExeFile           [260]byte
}

func main() {

	fmt.Println("正在守护pbbot")
	for {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
		pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
		if int(pHandle) == -1 {
			return
		}
		Process32Next := kernel32.NewProc("Process32Next")
		var pNameList []string
		for {
			var proc PROCESSENTRY32
			proc.dwSize = ulong(unsafe.Sizeof(proc))
			if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {
				pName := string(proc.szExeFile[0:])
				pNameList = append(pNameList, pName)
			} else {
				break
			}
		}

		if !Contain(pNameList, "pbbot") {
			exec.Command(`cmd`, `/c`, "pbbot.exe").Start()
			fmt.Println(time.Now().String(), "正在启动pbbot")
		}
		time.Sleep(time.Second * 30)
	}
}

func Contain(list []string, data string) bool {
	if data == "" || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if v == "" {
			continue
		}
		if strings.Contains(v, data) {
			return true
		}
	}
	return false
}
