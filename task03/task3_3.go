package main

import (
	"fmt"
	"sync"
)

func worker(numbers <-chan int, squares chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range numbers {
		squares <- number * number
	}
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	numbersCh := make(chan int, len(numbers))
	squaresCh := make(chan int, len(numbers))

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ { // Запускаем 2 горутины-работника
		wg.Add(1)
		go worker(numbersCh, squaresCh, &wg)
	}

	for _, number := range numbers {
		numbersCh <- number
	}
	close(numbersCh) // Закрытие канала, так как отправка чисел завершена

	wg.Wait()        // Ожидание завершения всех работников
	close(squaresCh) // Закрытие канала квадратов после завершения работников

	// Суммирование результатов
	sum := 0
	for square := range squaresCh {
		sum += square
	}

	fmt.Println("Sum of squares:", sum)
}
