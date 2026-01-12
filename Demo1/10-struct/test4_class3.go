package main

import "fmt"

// 多态
// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string // 获取动物颜色
	GetType() string  // 获取动物的种类
}

// 具体的类
type Cat struct {
	color string // 猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
}

func main() {
	var animal AnimalIF // 接口的数据类型，父类指针
	animal = &Cat{"Green"}
	//animal.Sleep()
	//fmt.Println(animal.GetColor())
	//fmt.Println(animal.GetType())
	showAnimal(animal)

	fmt.Println("--------------------------------")
	animal = &Dog{"Yellow"}
	//animal.Sleep()
	//fmt.Println(animal.GetColor())
	//fmt.Println(animal.GetType())
	showAnimal(animal)

	fmt.Println("--------------------------------")
	cat := Cat{"Green"}
	dog := Dog{"Yellow"}
	showAnimal(&cat)
	showAnimal(&dog)
}
