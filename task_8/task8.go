package task_8

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type sin struct {
	counter int64
	flag    uint64
	wait    *sync.Cond
}

func CripplesGroup() *sin {
	mut := &sync.Mutex{}
	return &sin{
		counter: 0,
		flag:    0,
		wait:    sync.NewCond(mut),
	}
}
func (wg *sin) Add(i int64) {
	wg.wait.L.Lock()
	defer wg.wait.L.Unlock()
	atomic.AddInt64(&wg.counter, i)
	atomic.AddUint64(&wg.flag, 1)
	wg.wait.Broadcast()
}

func (wg *sin) Done() {
	wg.wait.L.Lock()
	defer wg.wait.L.Unlock()
	atomic.AddInt64(&wg.counter, -1)
	wg.wait.Signal()
}

func (wg *sin) Wait() {
	wg.wait.L.Lock()
	defer wg.wait.L.Unlock()

	for {
		f := atomic.LoadUint64(&wg.flag)
		c := atomic.LoadInt64(&wg.counter)
		if f > 0 && c == 0 {
			break
		}
		wg.wait.Wait()
	}
}

var counter int64 = 0

func incrementCounter(wg *sin) {
	atomic.AddInt64(&counter, 1)
	wg.Done()
}

func StartTask8(count int) {
	fmt.Println("Start task 8")
	wg := CripplesGroup()

	for i := 0; i < count; i++ {
		wg.Add(1)
		go incrementCounter(wg)
	}

	wg.Wait()

	fmt.Println("data-in:", count, ", result increment:", counter)
}
