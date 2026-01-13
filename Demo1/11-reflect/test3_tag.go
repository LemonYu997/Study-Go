package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	// ``标签说明
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	// 获取全部元素
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		// 获取标签
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("tag info = ", tagInfo)
		fmt.Println("tag doc = ", tagDoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
