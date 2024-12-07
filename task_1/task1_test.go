package task_1

import "testing"

var tests = []struct {
	value interface{}
	res   string
}{
	{1, "int"},
	{_oct(02), "_oct"},
	{_hex(0x3), "_hex"},
	{3.14, "float64"},
	{"Golang", "string"},
	{true, "bool"},
	{complex64(complex(2, 4)), "complex64"},
	{complex128(complex(2, 4)), "complex128"},
}

func TestCheckType(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.res, func(t *testing.T) {
			ans := CheckType(tt.value)
			if ans != tt.res {
				t.Errorf("test result: %s, expected value: %s", ans, tt.res)
			}
		})
	}
}

var tests2 = []interface{}{42, _oct(02), _hex(0x3), 3.14, 1 + 2i, "Golang", true}

func TestConvertString(t *testing.T) {
	t.Run("tests convert array", func(t *testing.T) {
		ans := ConvertString(tests2)
		want := "42 O2 0x3 3.14 (1+2i) Golang true"
		if ans != want {
			t.Errorf("test result: %s, expected value: %s", ans, want)
		}
	})
}
