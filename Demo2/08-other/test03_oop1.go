package main

import "fmt"

/**
1、通过组合实现继承
2、通过接口实现多态

Speaker 是一个接口，定义了一个 Speak 方法。
Animal 结构体实现了 Speaker 接口。
Dog 结构体通过嵌入 Animal 结构体，间接实现了 Speaker 接口。
在 main 函数中，我们将 Dog 实例赋值给 Speaker 接口，并通过接口调用 Speak 方法。
*/
// 定义接口
type Speaker interface {
	Speak()
}

// 父结构体
type Animal struct {
	Name string
}

// 实现接口方法
func (a *Animal) Speak() {
	fmt.Println(a.Name, "says hello!")
}

// 子结构体
type Dog struct {
	Animal
	Breed string
}

func main() {
	var speaker Speaker

	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	// 通过组合实现继承
	dog.Speak() // 调用父结构体的方法
	fmt.Println("Breed:", dog.Breed)

	// 通过接口实现多态
	speaker = &dog
	speaker.Speak() // 通过接口调用方法
}
