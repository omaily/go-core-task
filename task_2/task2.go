package task_2

import (
	"fmt"
	"math/rand"
)

func createSlise(n int) []int {
	slise := make([]int, n)
	for i := range slise {
		slise[i] = rand.Intn(10)
	}
	return slise
}

func StartTask2() {
	// 2.1
	originalSlice := createSlise(10)
	fmt.Println("task 2:", originalSlice)
}
