package task_8

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

func incrementCounter() {
	counter++
	wg.Done()
}

func StartTask8(count int) {
	fmt.Println("Start task 8")

	for i := 0; i < count; i++ {
		wg.Add(1)
		go incrementCounter()
	}

	wg.Wait()
	fmt.Println("data-in:", count, ", result increment:", counter)
}
