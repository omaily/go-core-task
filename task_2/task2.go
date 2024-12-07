package task_2

import (
	"fmt"
	"math/rand"
)

func createSlise(n int) []int {
	sliсe := make([]int, n)
	for i := range sliсe {
		sliсe[i] = rand.Intn(10)
	}
	return sliсe
}

func sliceExample(sliсe []int) []int {
	even := make([]int, 0, len(sliсe))
	for _, v := range sliсe {
		if v%2 == 0 {
			even = append(even, v)
		}
	}

	return even
}

func StartTask2() {
	// 2.1
	originalSlice := createSlise(10)
	fmt.Println("task 2:", originalSlice)

	// 2.2
	evenSlice := sliceExample(originalSlice)
	fmt.Println("task 2:", evenSlice)

}
