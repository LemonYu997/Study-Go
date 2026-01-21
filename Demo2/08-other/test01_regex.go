package main

import (
	"fmt"
	"regexp"
)

// 正则表达式
func main() {
	//检查字符串是否匹配正则表达式
	pattern := `^[a-zA-Z0-9]+$`
	regex := regexp.MustCompile(pattern)

	str := "Hello123"
	if regex.MatchString(str) {
		fmt.Println("字符串匹配正则表达式")
	} else {
		fmt.Println("字符串不匹配正则表达式")
	}

	fmt.Println("---------------------------------------")

	// 查找匹配的字符串
	pattern = `\d+`
	regex = regexp.MustCompile(pattern)

	str = "我有 3 个苹果和 5 个香蕉"
	matches := regex.FindAllString(str, -1)
	fmt.Println("找到的数字：", matches)

	fmt.Println("---------------------------------------")

	//替换匹配的字符串
	pattern = `\s+`
	regex = regexp.MustCompile(pattern)

	str = "Hello    World"
	result := regex.ReplaceAllString(str, " ")
	fmt.Println("替换后的字符串：", result)

	fmt.Println("---------------------------------------")

	//分割字符串
	pattern = `,`
	regex = regexp.MustCompile(pattern)

	str = "apple,banana,orange"
	parts := regex.Split(str, -1)
	fmt.Println("分割后的字符串：", parts)
}
