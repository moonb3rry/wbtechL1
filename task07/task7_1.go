package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	m   map[int]int
	mux sync.Mutex
}

func (sm *SafeMap) Write(key, value int) {
	sm.mux.Lock() // Блокировка доступа к критической секции
	sm.m[key] = value
	sm.mux.Unlock() // Разблокировка доступа к критической секции
}

func main() {
	safeMap := SafeMap{m: make(map[int]int)}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Write(i, i*2)
		}(i)
	}

	wg.Wait() // Ожидание завершения всех горутин

	fmt.Println("map:", safeMap.m)
}
