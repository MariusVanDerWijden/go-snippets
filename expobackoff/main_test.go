package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestExponentialBackoff(t *testing.T) {
	var (
		n        = 10
		resource sync.Mutex
		timeSlot = 10 * time.Millisecond
		ctx, _   = context.WithTimeout(context.Background(), 10*time.Second)
		wg       = new(sync.WaitGroup)
	)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int, ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
				}
				backoff := 1
				ok := resource.TryLock()
				if ok {
					fmt.Printf("Thread %v is doing some work\n", i)
					time.Sleep(timeSlot)
				} else {
					fmt.Printf("Thread %v is backing off for %v slots\n", i, backoff)
					backoff = backoff << 1
					time.Sleep(time.Duration(backoff))
				}
			}
		}(i, ctx)
	}
	wg.Wait()
}
