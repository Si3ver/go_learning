package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		a, b = b, a+b
	}
}

// 自动类型推断
// 多变量赋值
// 递增声明变量
