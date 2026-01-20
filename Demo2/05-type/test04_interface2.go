package main

import "fmt"

type Reader2 interface {
	Read() string
}

type Writer2 interface {
	Write(data string)
}

// 接口嵌套组合
type ReadWriter interface {
	Reader2
	Writer2
}

type File struct{}

func (f File) Read() string {
	return "Reading data"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}

func main() {
	var rw ReadWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("Hello, Go!")
}
