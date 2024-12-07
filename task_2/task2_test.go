package task_2

import (
	"slices"
	"testing"
)

var testsData = []struct {
	name string
	data []int
}{
	{"origin", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{"not even", []int{1, 3, 5, 7, 9, 11, 13, 15}},
	{"even", []int{2, 4, 6, 8, 10, 12, 14, 16}},
	{"unit", []int{1}},
	{"empty", []int{}},
	{"nil", nil},
}

func TestSliceExample(t *testing.T) {
	for _, tt := range testsData {
		t.Run(tt.name, func(t *testing.T) {
			ans := SliceExample(tt.data)
			for _, v := range ans {
				if v%2 != 0 {
					t.Errorf("there are odd numbers in the set")
				}
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	for _, tt := range testsData {
		t.Run(tt.name, func(t *testing.T) {
			ans := AddElements(tt.data, 7)
			if len(tt.data)+1 != len(ans) {
				t.Errorf("the array did not become larger. was: %d, become: %d", len(tt.data), len(ans))
			}
		})
	}
}

func TestCopyElements(t *testing.T) {
	for _, tt := range testsData {
		t.Run(tt.name, func(t *testing.T) {
			ans := CopyElements(tt.data)
			if !slices.Equal(ans, tt.data) {
				t.Errorf("different values of cuts: originalSlice: %v, copySlice %v", tt.data, ans)
			}
			if len(tt.data) > 0 && &tt.data[0] == &ans[0] {
				t.Error("return was not a copy but the same cut", tt.data)
			}
		})
	}
}
