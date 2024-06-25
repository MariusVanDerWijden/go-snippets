package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		n          = 10
		resource   sync.Mutex
		timeSlot   = 10 * time.Millisecond
		ctx, _     = context.WithTimeout(context.Background(), 10*time.Second)
		wg         = new(sync.WaitGroup)
		start      = time.Now()
		collisions = new(atomic.Int32)
	)

	expRandBackoff := func(i int, backoff int) int {
		rng := rand.Int31n(int32(backoff))
		fmt.Printf("Thread %v is backing off for %v slots, max %v\n", i, rng, backoff)
		time.Sleep(time.Duration(rng) * timeSlot)
		return backoff << 1
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int, ctx context.Context, backoffStrat func(int, int) int) {
			defer wg.Done()
			backoff := 1
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				ok := resource.TryLock()
				if ok {
					fmt.Printf("Thread %v is doing some work\n", i)
					time.Sleep(timeSlot)
					resource.Unlock()
					return
				} else {
					collisions.Add(1)
					backoff = backoffStrat(i, backoff)
				}
			}
		}(i, ctx, expRandBackoff)
	}
	wg.Wait()
	fmt.Printf("Time spent to finish: %v, optimal: %v, collisions: %v\n", time.Since(start), time.Duration(n)*timeSlot, collisions.Load())
}
