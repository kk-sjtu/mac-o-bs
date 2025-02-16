package main

import "fmt"

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

type geometry interface {
	area() float64
	perim() float64
}

// 接口标准化了两个结构体。rect 和 circle 都实现了 geometry 接口，因此可以在 measure 函数中使用它们。
// 一个结构体实现了两个图形的这两个函数的功能。
// 任何实现了这两个函数area()，perim()的结构体都可以被称为 geometry 类型。

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	switch g.(type) {
	case circle:
		fmt.Println("It's a circle")
	default:
		fmt.Println("It's not a circle")
	}
}
