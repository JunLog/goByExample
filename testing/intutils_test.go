package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 名称以 Test 开头的函数来创建测试
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2,-2） = %d;want -2", ans)
		// t.Error,报告失败信息然后继续测试
		//  t.Fail 报告失败信息并立即终止
	}
}

// 名称以 Test 开头的函数来创建测试
func TestIntMinTableDriven(t *testing.T) {
	// 单元测试可重复，常用 表驱动 风格编写调用测试
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// t.Run 可以运行一个 subtests 子测试，一个子测试对应一行数据
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d,want %d", ans, tt.want)
			}
		})
	}
}

// 运行命令 ： go test -v
