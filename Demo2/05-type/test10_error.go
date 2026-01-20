package main

import (
	"errors"
	"fmt"
)

// 自定义错误
type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

// 显式返回错误
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 自定义错误
func divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

// 错误处理
func main() {
	err := errors.New("this is an error")
	fmt.Println(err) // 输出：this is an error

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	_, err2 := divide2(10, 0)
	if err2 != nil {
		fmt.Println(err2) // 输出：cannot divide 10 by 0
	}
}
