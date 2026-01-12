package main

import "fmt"

// 封装
// 当类名首字母大写时，表示其他包可以访问
type Hero struct {
	// 当属性名首字母大写时，表示其他包可以访问
	Name  string
	Ad    int
	Level int
}

/*func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) GetName() string {
	// this 是调用该方法对象的一个拷贝
	fmt.Println("Name = " + this.Name)
	return this.Name
}

func (this Hero) SetName(newName string) {
	this.Name = newName
}*/

// 当方法名首字母大写时，表示其他包可以访问
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() string {
	// this 是调用该方法对象的一个拷贝
	fmt.Println("Name = " + this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "zhangsan", Ad: 100, Level: 1}

	hero.Show()
	fmt.Println("----------------------------------------")

	hero.SetName("li4")
	hero.Show()
}
