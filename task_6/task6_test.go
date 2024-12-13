package task_6

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestStartTask6(t *testing.T) {
	t.Run("testing generator", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Recovered: %v", r)
			}
		}()

		timeout := make(chan struct{})
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
		defer cancel()
		go func() {
			time.Sleep(time.Millisecond * 100)
			StartTask6(ctx)
			timeout <- struct{}{}
		}()

		select {
		case <-timeout:
			fmt.Println("ok")
		case <-ctx.Done():
			t.Errorf("context deadline exceeded")
		}

	})
}
