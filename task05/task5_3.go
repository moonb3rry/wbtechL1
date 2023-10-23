package main

import (
	"fmt"
	"time"
)

// ограничение по количеству сообщений
func main() {
	ch := make(chan int)
	duration := 5 * time.Second
	endTime := time.Now().Add(duration) // Вычисляем время окончания

	go func() {
		for i := 0; time.Now().Before(endTime); i++ { // Проверка текущего времени
			ch <- i
		}
		close(ch) // Закрываем канал, когда время истекло
	}()

	for value := range ch {
		fmt.Println(value)
	}

	fmt.Println("Время истекло")
}
