package task_6

import (
	"context"
	"fmt"
	"math/rand/v2"
)

func StartTask6(ctx context.Context) {
	ch := make(chan int)
	go func(c chan<- int) {
		c <- rand.IntN(100)
	}(ch)

	fmt.Println("генерируем случайное число:", <-ch)
}
