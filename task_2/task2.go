package task_2

import (
	"fmt"
	"math/rand"
)

func createSlise(n int) []int {
	slice := make([]int, n, 15)
	for i := range slice {
		slice[i] = rand.Intn(10)
	}
	return slice
}

func sliceExample(slice []int) []int {
	even := make([]int, 0, len(slice))
	for _, v := range slice {
		if v%2 == 0 {
			even = append(even, v)
		}
	}

	return even
}

func addElements(slece []int, v int) []int {
	slece2 := make([]int, len(slece), len(slece)+1)
	copy(slece2, slece)
	slece2 = append(slece2, v)
	return slece2
}

func StartTask2() {
	// 2.1
	originalSlice := createSlise(10)
	fmt.Println("task 2.1:", originalSlice)

	// 2.2
	evenSlice := sliceExample(originalSlice)
	fmt.Println("task 2.2:", evenSlice)

	// 2.3
	addSlice := addElements(originalSlice, 12)
	fmt.Println("task 2.3:", addSlice)
}
