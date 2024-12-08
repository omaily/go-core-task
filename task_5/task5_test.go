package task_5

import (
	"reflect"
	"testing"
)

var testsData = []struct {
	name  string
	data1 []int
	data2 []int
	res   []int
}{
	{
		"origin",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}, {
		"no matches",
		[]int{0, 1, 2, 3, 4},
		[]int{5, 6, 7, 8, 9},
		[]int{},
	}, {
		"same",
		[]int{1, 1, 1, 2, 2, 2, 3, 3, 3},
		[]int{1, 1, 2, 2, 3, 3, 4, 4},
		[]int{1, 1, 2, 2, 3, 3},
	}, {
		"repetit Slise1",
		[]int{1, 1, 1, 2, 2, 2, 3, 3, 3},
		[]int{1, 2, 3, 4, 1, 2, 3, 4},
		[]int{1, 2, 3, 1, 2, 3},
	}, {
		"empty",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{},
		[]int{},
	}, {
		"null",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		nil,
		[]int{},
	},
}

func TestIntersect(t *testing.T) {
	for _, tt := range testsData {
		t.Run(tt.name, func(t *testing.T) {
			_, ans := Intersect(tt.data1, tt.data2)
			if !reflect.DeepEqual(ans, tt.res) {
				t.Errorf("was expected: %v, become: %v", tt.res, ans)
			}
		})
	}
}
