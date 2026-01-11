package main

import (
	"Study-Go/Demo1/5-init/lib1"

	// 导入但不使用时 前边使用 _ (表示匿名） 可以防止报错
	//_ "Study-Go/Demo1/5-init/lib1"
	// 可以定义包别名
	//myLib2 "Study-Go/Demo1/5-init/lib2"
	// 可以使用 . 作为别名，可以直接调用包中的方法 不推荐使用，防止同名方法无法读懂
	. "Study-Go/Demo1/5-init/lib2"
)

func main() {
	lib1.Lib1Test()

	//myLib2.Lib2Test()
	Lib2Test()
}
