package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {

	// 图片大小
	const size = 300
	// 创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	//从0 到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0——2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		//用黑色回执sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
		file, err := os.Create("test.png")
		// 判断是否异常
		if err != nil {
			log.Fatal(err)
		}
		png.Encode(file, pic)
		file.Close()
	}

}
