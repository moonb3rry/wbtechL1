package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5} // Наш массив чисел

	// Создание каналов
	numberChan := make(chan int)
	resultChan := make(chan int)

	// Горутина для умножения числа на 2
	go func() {
		for number := range numberChan {
			resultChan <- number * 2
		}
		close(resultChan) // Закрываем resultChan после обработки всех чисел
	}()

	// Отправка чисел в канал
	go func() {
		for _, number := range numbers {
			numberChan <- number
		}
		close(numberChan) // Закрываем numberChan, когда все числа отправлены
	}()

	// Вывод результатов
	for result := range resultChan {
		fmt.Println(result)
	}
}
