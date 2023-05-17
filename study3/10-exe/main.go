package main

import (
	"fmt"
	"os/exec"
)

// start D:\wdrcode\gq22156\gq22156.exe
func main() {
	// 创建 cmd 对象
	cmd := exec.Command("cmd", "/c", "start D:\\app\\gq22156.exe")

	// 执行命令并获取输出结果
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Command execution failed:", err)
		return
	}

}
