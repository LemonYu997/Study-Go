package main

import "fmt"

// 基础语法
func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	// Sprintf返回格式化字符串
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	// Println执行标准输出
	fmt.Println(target_url)
}
