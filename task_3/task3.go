package task_3

import (
	"fmt"
)

type StringIntMap map[string]int

func (sm StringIntMap) Add(key string, value int) {
	sm[key] = value
}

func (sm StringIntMap) Remove(key string) {
	delete(sm, key)
}

func (sm StringIntMap) Copy() map[string]int {
	m := make(map[string]int)
	for key := range sm {
		m[key] = sm[key]
	}
	return m
}

func (sm StringIntMap) Exists(key string) bool {
	_, ok := sm[key]
	return ok
}

func (sm StringIntMap) Get(key string) (int, bool) {
	v, ok := sm[key]
	return v, ok

}

func StartTask3() {
	m := StringIntMap{
		"age":    27,
		"height": 187,
		"waight": 84,
	}

	// 3.1
	m.Add("mmr", 3500)
	fmt.Println("task 3.1:", m)

	// 3.2
	m.Remove("age")
	fmt.Println("task 3.2:", m)

	// 3.3
	m2 := m.Copy()
	fmt.Println("task 3.3:", m2)

	// 3.4
	isExists := m.Exists("age")
	fmt.Println("task 3.4:", isExists)

	// 3.5
	value, ok := m.Get("mmr")
	fmt.Println("task 3.4:", value, ok)
}
