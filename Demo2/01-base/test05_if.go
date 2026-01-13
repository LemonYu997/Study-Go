package main

import (
	"fmt"
	"time"
)

func main() {
	a := 10

	// if else
	if a < 20 {
		fmt.Println("a < 20")
	} else {
		fmt.Println("a >= 20")
	}
	fmt.Println("a = ", a)

	b := 100
	c := 100

	// 嵌套 if
	if b == 100 {
		if c == 200 {
			fmt.Println("b == 100, c = 200")
		}
	}

	// switch
	grade := "B"
	marks := 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B", grade == " C":
		fmt.Println("良好")
	case grade == "D":
		fmt.Println("及格")
	case grade == "F":
		fmt.Println("不及格")
	default:
		fmt.Println("差")
	}

	// 类型判断
	var x interface{}
	fmt.Printf("x type: %T\n", x)
	switch i := x.(type) {
	case nil:
		fmt.Printf("x is nil\n")
	case int:
		fmt.Printf("x is int\n", i)
	case float64:
		fmt.Printf("x is float64\n", i)
	case func(int) float64:
		fmt.Printf("x type: %T\n", i)
	case string, bool:
		fmt.Printf("x is string or bool\n", i)
	default:
		fmt.Printf("未知型")
	}

	// fallthrough 会强制执行后面的 case 语句, 不会判断下一条 case 的表达式结果是否为 true
	//y := true
	switch {
	// 这里是编译器报错，不影响运行
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	// 执行结果
	//2、case 条件语句为 true
	//3、case 条件语句为 false
	//4、case 条件语句为 true

	// select 判断
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
}
