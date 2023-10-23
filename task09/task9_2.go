package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5} // Наш массив чисел

	numberChan := make(chan int)
	resultChan := make(chan int)

	var wg sync.WaitGroup

	// Создание нескольких горутин для обработки умножения
	for i := 0; i < 3; i++ { // Например, создаем 3 горутины
		wg.Add(1)
		go func() {
			for number := range numberChan {
				resultChan <- number * 2
			}
			wg.Done()
		}()
	}

	// Отправка чисел в канал
	go func() {
		for _, number := range numbers {
			numberChan <- number
		}
		close(numberChan) // Закрываем numberChan, когда все числа отправлены
	}()

	// Закрыть resultChan после завершения всех горутин умножения
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Вывод результатов
	for result := range resultChan {
		fmt.Println(result)
	}
}
