package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	// 首字母必须大写，模板文件才能正确读取到
	Name   string
	Gender string
	Age    int
}

// template 模版引擎
func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1、定义模版 hello.tmpl
	// 2、解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v", err)
		return
	}
	// 3、渲染模版
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name":   "小公主",
		"gender": "女",
		"age":    15,
	}
	hobbyList := []string{
		"篮球", "足球", "双色球",
	}
	err = t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("Execute template failed, err:%v", err)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
