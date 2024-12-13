package task_7

import (
	"fmt"
	"sync"
)

func joinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(len(chs))
	go func() {
		for _, ch := range chs {
			go func() {
				for id := range ch {
					mergedCh <- id
				}
				wg.Done()
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

func StartTask7() {

	a := make(chan int)
	b := make(chan int)

	go func() {
		for v := range 5 {
			a <- v + 1
		}
		close(a)
	}()

	go func() {
		for v := range 5 {
			a <- (v + 1) * 10
		}
		close(b)
	}()

	for num := range joinChannels(a, b) {
		fmt.Println(num)
	}

}
