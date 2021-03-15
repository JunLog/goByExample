package main

import "fmt"

type person struct { // 结构体 是 可变（mutable）
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "张三", age: 40}) //结构体指针

	s := person{name: "李四", age: 50}
	fmt.Println(s.name)

	sp := &s            // 取结构体指针
	fmt.Println(sp.age) // 直接访问字段（能自动解引用）

	sp.age = 60
	fmt.Println(sp.age)
}
