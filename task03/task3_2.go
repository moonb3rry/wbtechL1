package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n

			mu.Lock()
			sum += square
			mu.Unlock()
		}(number)
	}

	wg.Wait() // Ожидание завершения всех горутин
	fmt.Println("Sum of squares:", sum)
}
