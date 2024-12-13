package task_9

import (
	"fmt"
	"time"
)

func square(ch chan<- float64, base int8) {
	base64 := float64(base)
	extent := base64 * base64 * base64
	ch <- extent
}

func StartTask9() {
	short := make(chan int8)
	long := make(chan float64)

	go func() {
		defer close(short)
		for i := range int8(8) {
			time.Sleep(100 * time.Millisecond)
			short <- i
		}
	}()

	go func() {
		for {
			c, ok := <-short
			if !ok {
				close(long)
				break
			}
			square(long, c)
		}
	}()

	for c := range long {
		fmt.Println(c)
	}
}
