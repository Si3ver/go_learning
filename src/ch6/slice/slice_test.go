package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2)) // length代表可访问的元素个数，capacity代表容量
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2, len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(s, len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // 3, 9
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer)) // 3, 7
	summer[0] = "Unknow"
	t.Log(Q2, summer, year) // 共享存储!!!
}

func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// if a == b { // 切片不可以比较
	// 	t.Log("equal")
	// } else {
	// 	t.Log("not equal!")
	// }
	t.Log(a, b)
}
