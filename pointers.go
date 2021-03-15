package main

import "fmt"

func zeroval(ival int) { //值传递
	ival = 0
}

func zeroptr(iptr *int) { //引用传递，参数的 *int 表示使用 int指针
	*iptr = 0 // *iptr 是指解引用，从内存中获取对应的值
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) // 改变不了值
	fmt.Println("zeroval:", i)

	zeroptr(&i) // &i 取得 i 的内存地址（指向 i 的指针）
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
