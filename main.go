package main

import (
	task "coretask/task_4"
	"fmt"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	res := task.StartTask4(slice1, slice2)
	fmt.Println(res)
}
