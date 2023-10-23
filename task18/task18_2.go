package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter2 struct {
	value int64
}

func (c *Counter2) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter2) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := Counter2{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())
}
