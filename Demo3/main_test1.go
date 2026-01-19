package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(writer http.ResponseWriter, request *http.Request) {
	//_, _ = fmt.Fprintln(writer, "<h1>Hello World!</h1><h2>你好</h2>")

	//从文件中读取
	b, _ := os.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(writer, string(b))
}

func main() {
	http.HandleFunc("/hello", sayHello)

	// 监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
