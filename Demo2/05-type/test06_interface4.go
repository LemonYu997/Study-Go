package main

import "fmt"

type Shape2 interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle2 struct {
	radius float64
}

func (c Circle2) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape2

	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积: %f\n", s.area())

	s = Circle2{radius: 3}
	fmt.Printf("圆形面积: %f\n", s.area())
}
