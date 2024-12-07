package task_4

import (
	"reflect"
	"testing"
)

var testsData = []struct {
	name  string
	data1 []string
	data2 []string
	res   []string
}{
	{
		"origin",
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		[]string{"banana", "date", "fig"},
		[]string{"apple", "cherry", "43", "lead", "gno1"},
	}, {
		"same",
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		[]string{},
	}, {
		"repetit Slise1",
		[]string{"apple", "apple", "apple", "apple", "apple", "banana", "fig"},
		[]string{"banana", "date", "fig"},
		[]string{"apple", "apple", "apple", "apple", "apple"},
	}, {
		"repetit Slise2",
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		[]string{"banana", "banana", "cherry", "cherry", "date", "date"},
		[]string{"apple", "43", "lead", "gno1"},
	}, {
		"empty",
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		[]string{},
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
	}, {
		"null",
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		nil,
		[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
	},
}

func TestStartTask4(t *testing.T) {
	for _, tt := range testsData {
		t.Run(tt.name, func(t *testing.T) {
			ans := StartTask4(tt.data1, tt.data2)
			if !reflect.DeepEqual(ans, tt.res) {
				t.Errorf("was expected: %v, become: %v", tt.res, ans)
			}
		})
	}
}
