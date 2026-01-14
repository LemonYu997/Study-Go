package main

import "fmt"

func displayImage(img [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%3d ", img[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// 模拟一个简单的3x3灰度图像
	// 每个像素值范围：0(黑色) ~ 255(白色)
	var image [3][3]int

	// 初始化图像（一个简单的渐变）
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			image[i][j] = (i + j) * 50
		}
	}

	// 显示原始图像
	fmt.Println("原始图像:")
	displayImage(image)

	// 图像处理：增加亮度
	fmt.Println("\n增加亮度后的图像:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// 增加亮度，但不超过255
			newValue := image[i][j] + 50
			if newValue > 255 {
				newValue = 255
			}
			image[i][j] = newValue
		}
	}
	displayImage(image)
}
