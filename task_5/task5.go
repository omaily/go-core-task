package task_5

import "fmt"

func Intersect(slice1 []int, slice2 []int) (bool, []int) {
	if len(slice1) < len(slice2) {
		return Intersect(slice2, slice1)
	}
	var flag bool = false
	m := make(map[int]int, len(slice1))
	for _, key := range slice1 {
		m[key]++
	}

	intersection := make([]int, 0, len(slice2))
	for _, key := range slice2 {
		if m[key] > 0 {
			flag = true
			intersection = append(intersection, key)
			m[key]--
		}
	}

	return flag, intersection
}

func StartTask5() {
	// уловия когда в слайсе находятся повторающиеся значения
	// в своем примере каждый элемент выходного слайса встречаться столько раз, сколько он отображается в обоих массивах
	a := []int{65, 3, 3, 3, 58, 678, 64}
	b := []int{64, 1, 3, 3, 43}
	fmt.Println(Intersect(a, b))
}
