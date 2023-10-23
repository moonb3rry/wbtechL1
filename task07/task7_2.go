package main

import (
	"fmt"
	"sync"
)

func main() {
	var syncMap sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncMap.Store(i, i*2) // Store используется для безопасной записи пары ключ/значение
		}(i)
	}

	wg.Wait()

	syncMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
}
