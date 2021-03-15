package main

import (
	"fmt"
	"math"
)

// 接口：对方法集进行命名和分组的一种机制

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 { // 方法
	return r.width * r.height
}
func (r rect) perim() float64 { // 方法
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 { // 方法
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 { // 方法
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { //接口 （通用函数）
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r) // r 和 c 都实现了 geometry 接口
	measure(c)
}
