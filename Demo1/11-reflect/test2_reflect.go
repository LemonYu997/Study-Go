package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

type User struct {
	Id   int
	Name string
	Age  int
}

// 1,17之后定义方法不能加*才能被反射获取
// func (this *User) Call() {
func (this User) Call() {
	fmt.Println("user is called .. ")
	fmt.Printf("%v\n", this)
}

func main() {
	var num float64 = 1.2345

	reflectNum(num)
	fmt.Println("------------------------------------")

	user := User{1, "Aceld", 18}
	user.Call()
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("type : ", inputType)
	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("value : ", inputValue)

	// 通过type 获取里面的字段
	// 1、获取interface的reflect.Type 得到NumField，进行遍历
	// 2、得到每个field，数据类型
	// 3、通过field有一个Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	//通过type 获取里面的方法、调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
