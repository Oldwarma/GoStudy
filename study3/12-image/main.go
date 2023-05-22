package main

import (
	"image"
	"image/png"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100)) // 创建一个大小为 100x100 的 RGBA 图像
	file, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img) // 将图像编码为 PNG 格式，并保存到文件中
	if err != nil {
		panic(err)
	}
}
