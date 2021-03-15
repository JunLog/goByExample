package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}

// fact(7)
// 7*fact(6)
// 7*6*fact(5)
// 7*6*5*fact(4)
// 7*6*5*4*fact(3)
// 7*6*5*4*3*fact(2)
// 7*6*5*4*3*2*fact(1)
// 7*6*5*4*3*2*1

// 这过程不断的使用栈保存和恢复，递归过大的话，可能会造成栈溢出

//尾递归-----------------------------------------------------------
// 尾部递归是指递归函数在调用自身后直接传回其值，而不对其再加运算，效率将会极大的提高。

// 如果一个函数中所有递归形式的调用都出现在函数的末尾，我们称这个递归函数是尾递归的。当递归调用是整个函数体中最后执行的语句且它的返回值不属于表达式的一部分时，这个递归调用就是尾递归。尾递归函数的特点是在回归过程中不用做任何操作，这个特性很重要，因为大多数现代的编译器会利用这种特点自动生成优化的代码。-- 来自百度百科。

// 尾递归函数，部分高级语言编译器会进行优化，减少不必要的堆栈生成，使得程序栈维持固定的层数，不会出现栈溢出的情况。

// package main
// import "fmt"

// func RescuvieTail(n int, a int) int {
//     if n == 1 {
//         return a
//     }

//     return RescuvieTail(n-1, a*n)
// }

// func main() {
//     fmt.Println(RescuvieTail(5, 1))
// }

// RescuvieTail(5, 1)
// RescuvieTail(4, 1*5)=RescuvieTail(4, 5)
// RescuvieTail(3, 5*4)=RescuvieTail(3, 20)
// RescuvieTail(2, 20*3)=RescuvieTail(2, 60)
// RescuvieTail(1, 60*2)=RescuvieTail(1, 120)
// 120
