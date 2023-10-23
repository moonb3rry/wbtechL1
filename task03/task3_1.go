package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	squaresCh := make(chan int) // Канал для квадратов чисел

	// Запуск горутин для вычисления квадратов
	for _, number := range numbers {
		go func(n int, ch chan<- int) {
			ch <- n * n
		}(number, squaresCh)
	}

	// Суммирование результатов
	sum := 0
	for range numbers {
		sum += <-squaresCh
	}

	fmt.Println("Sum of squares:", sum)
}
