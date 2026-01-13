package main

import (
	"fmt"
	"io"
	"os"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("ReadBook")
}

func (this *Book) WriteBook() {
	fmt.Println("WriteBook")
}

func main() {
	var a string
	// pair<statictype:string, value:"aceld">
	a = "aceld"

	// 万能类型
	var allType interface{}
	// pair<type:string, value:"aceld">
	allType = a
	// 用断言获取其值
	str, _ := allType.(string)
	fmt.Println(str) //"aceld"

	//tty: pair<type:*os.File, value: "/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

	var r io.Reader
	// r: pair<type:*os.File, value:"/dev/tty"文件描述符>
	r = tty

	var w io.Writer
	// w: pair<type:*os.File, value:"/dev/tty"文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("hello world"))

	fmt.Println("------------------------------")

	// b: pair<type:Book, value:book{}地址>
	b := &Book{}

	var r1 Reader
	// r1: pair<type:Book, value:book{}地址>
	r1 = b

	r1.ReadBook()

	var w1 Writer
	// w1: pair<type:Book, value:book{}地址>
	w1 = r1.(Writer) // 此处的断言为什么能成功，因为w r具体的type是一致的
	w1.WriteBook()
}
