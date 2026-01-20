package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findItem(id int) error {
	return fmt.Errorf("database error: %w", ErrNotFound)
}

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func getError() error {
	return &MyError{Code: 404, Msg: "Not Found"}
}

// panic 和 recover
//
//	panic 用于处理不可恢复的错误，recover 用于从 panic 中恢复。
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went wrong")
}

func main() {
	err := findItem(1)

	//errors.Is
	//检查某个错误是否是特定错误或由该错误包装而成。
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Item not found")
	} else {
		fmt.Println("Other error:", err)
	}

	err = getError()
	var myErr *MyError
	// errors.As
	// 将错误转换为特定类型以便进一步处理。
	if errors.As(err, &myErr) {
		fmt.Printf("Custom error - Code: %d, Msg: %s\n", myErr.Code, myErr.Msg)
	}

	fmt.Println("Starting program...")
	safeFunction()
	fmt.Println("Program continued after panic")
}
