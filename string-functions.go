package main

import (
	"fmt"
	s "strings" // 别名
	/*
		import
		import "fmt"最常用的一种形式

		import "./test"导入同一目录下test包中的内容

		import f "fmt"导入fmt，并给它启别名 ｆ

		import . "fmt"，将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加 fmt. ，如fmt.Println可以直接写成Println

		import _ "fmt" 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容。注意：这种形式的import，当import时就执行了fmt包中的init函数，而不能够使用该包的其他函数。
	*/)

var p = fmt.Println // 别名

func main() {
	p("Contains:	", s.Contains("test", "es"))      // 是否含有
	p("Count:		", s.Count("test", "t"))            // 统计数量
	p("HasPrefix:	", s.HasPrefix("test", "te"))    // 开头
	p("HsaSuffix:	", s.HasSuffix("test", "st"))    // 结尾
	p("Index:		", s.Index("test", "e"))            // 索引
	p("Join:		", s.Join([]string{"a", "b"}, "-"))  // 拼接
	p("Repeat:		", s.Repeat("a", 5))               // 重复
	p("Replace:	", s.Replace("foo", "o", "0", -1)) // 替换所有
	p("Replace:	", s.Replace("foo", "o", "0", 1))  // 替换 1 次
	p("Split:		", s.Split("a-b-c-d-e", "-"))       // 分割
	p("ToLower:	", s.ToLower("TEST"))
	p("ToUpper:	", s.ToUpper("test"))
	p()

	p("Len:		", len("hello"))
	p("Char:		", "hello"[1]) //101 , "e" 的 utf8 10进制编码为 101（16进制为65）
}
