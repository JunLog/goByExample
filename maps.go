package main

import "fmt"

func main() {
	m := make(map[string]int) //make( map[键类型]值类型 )

	m["k1"] = 7 //设置键值对
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2") // 删除一个键值对
	fmt.Println("map:", m)

	_, prs := m["k2"] // _ 下划线（空白标识符）
	//键值对取值时，可以接收两个结果
	//一个是值（可以是空值或零值），一个是键是否存在的布尔值
	// 这里不需要这个值，可以使用 空白标识符 接收
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // 声明并初始化map（键值对）
	fmt.Println("map:", n)
}
