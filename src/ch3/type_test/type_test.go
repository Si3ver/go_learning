package type_test

import (
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	var c MyInt
	b = int64(a)
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") // 空字符串
	t.Log(len(s))
}
