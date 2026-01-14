package main

import "fmt"

// 学生成绩管理系统
func main() {
	// 定义：3个学生，每个学生有4门课程
	var grades [3][4]float64

	// 输入学生成绩
	grades = [3][4]float64{
		{85.5, 90.0, 78.5, 92.0}, // 学生1的成绩
		{88.0, 76.5, 95.0, 89.5}, // 学生2的成绩
		{92.5, 85.0, 88.5, 90.0}, // 学生3的成绩
	}

	// 计算每个学生的平均分
	fmt.Println("学生成绩统计:")
	for i, studentGrades := range grades {
		sum := 0.0
		for _, grade := range studentGrades {
			sum += grade
		}
		average := sum / float64(len(studentGrades))
		fmt.Printf("学生 %d: 平均分 = %.2f\n", i+1, average)
	}

	// 计算每门课程的平均分
	fmt.Println("\n课程平均分:")
	for j := 0; j < 4; j++ {
		sum := 0.0
		for i := 0; i < 3; i++ {
			sum += grades[i][j]
		}
		average := sum / 3.0
		fmt.Printf("课程 %d: 平均分 = %.2f\n", j+1, average)
	}
}
