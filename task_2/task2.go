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

func SliceExample(slice []int) []int {
	even := make([]int, 0, len(slice))
	for _, v := range slice {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return even
}

func AddElements(slice []int, v int) []int {
	slice2 := make([]int, len(slice), len(slice)+1)
	copy(slice2, slice)
	slice2 = append(slice2, v)
	return slice2
}

func CopyElements(slice []int) []int {
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	return slice2
}

func RemoveElement(slice []int, i int) []int {
	if i >= len(slice) || i < 0 {
		return slice
	}

	tempSlice := make([]int, i, len(slice))
	for i := range i {
		tempSlice[i] = slice[i]
	}

	return append(tempSlice, slice[i+1:]...)
}

func StartTask2() {
	// 2.1
	originalSlice := createSlise(10)
	fmt.Println("task 2.1:", originalSlice)

	// 2.2
	evenSlice := SliceExample(originalSlice)
	fmt.Println("task 2.2:", evenSlice)

	// 2.3
	addSlice := AddElements(originalSlice, 12)
	fmt.Println("task 2.3:", addSlice)

	// 2.4
	copySlice := CopyElements(originalSlice)
	originalSlice[1] = 21
	fmt.Println("task 2.4: originalSlice", originalSlice)
	fmt.Println("              copySlice", copySlice)

	// 2.5
	removeSlice := RemoveElement(originalSlice, 0)
	fmt.Println("task 2.5: originalSlice", originalSlice)
	fmt.Println("            removeSlice", removeSlice)
}
