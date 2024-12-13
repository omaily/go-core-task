package main

import (
	"context"
	task2 "coretask/task_2"
	task3 "coretask/task_3"
	task4 "coretask/task_4"
	task5 "coretask/task_5"
	task6 "coretask/task_6"
	"fmt"
)

func main() {
	task2.StartTask2()

	task3.StartTask3()

	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	res := task4.StartTask4(slice1, slice2)
	fmt.Println(res)

	task5.StartTask5()

	task6.StartTask6(context.Background())
}
