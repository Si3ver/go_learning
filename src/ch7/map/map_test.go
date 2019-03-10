package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9} // 方式一、声明并初始化
	t.Log(m1, m1[2], len(m1))
	m2 := map[int]int{} // 方式二、声明空map
	m2[4] = 16
	t.Log(m2, m2[2], len(m2))
	m3 := make(map[int]int, 10) // 方式三、利用map，并初始化一个容量（提高性能）
	t.Log(m3, m3[2], len(m3))
}

func TestAccessNotExitingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 0
	m1[2] = 0
	t.Log(m1[2]) // 0
	// m1[3] = 0
	if v, ok := m1[3]; ok { // ！！！判断map中某元素是否存在！！！
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 { // ！！！遍历！！！
		t.Log(k, v)
	}
}
