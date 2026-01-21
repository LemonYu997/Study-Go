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
		fmt.Println("Execute template failed, err:%v\n", err)
	}
}

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数 kua，要么只有一个返回值，要么有两个返回值，第二个必须是error
	k := func(name string) (string, error) {
		return name + "年轻又帅气！", nil
	}
	// 解析模版
	t := template.New("f1.tmpl") // 创建一个名字是f1的模版对象，一定要与模版的名字对上
	// 告诉模版引擎，现在多了一个自定义的函数 kua
	t.Funcs(template.FuncMap{
		"kua99": k,
	})
	_, err := t.ParseFiles("./f1.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
	}
	name := "小王子"
	// 渲染模版
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
		return
	}
	msg := "小王子"
	// 渲染模版
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
		return
	}
	msg := "小王子"
	// 渲染模版
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模版（模版继承方式）
	// 解析模版 根模板要写在前边
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
		return
	}
	msg := "小王子"
	// 渲染模版
	t.Execute(w, msg)
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模版（模版继承方式）
	// 解析模版 根模板要写在前边
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
		return
	}
	msg := "小王子"
	// 渲染模版
	t.Execute(w, msg)
}

func index3(w http.ResponseWriter, r *http.Request) {
	// 定义模版（模版继承方式）
	// 解析模版 修改模版引擎的标识符
	t, err := template.New("index3.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index3.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
		return
	}
	msg := "小王子"
	// 渲染模版
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Execute template failed, err:%v\n", err)
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模版（模版继承方式）
	// 解析模版 定义一个自定义的函数safe
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("Parse template failed, err:%v\n", err)
		return
	}
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='www.baidu.com'>百度</a>"
	// 渲染模版
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Println("Execute template failed, err:%v\n", err)
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/f1", f1)
	http.HandleFunc("/demo1", demo1)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)

	http.HandleFunc("/index3", index3)
	http.HandleFunc("/xss", xss)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
