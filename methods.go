package main

import (
	"fmt"
)

type rect struct {
	width, height int
}
// 为结构体定义方法，有点类似于类方法
func (r *rect) area() int { //指针类型接受者
	return r.width * r.height
}

func (r rect) perim() int { // 值类型接受者
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
