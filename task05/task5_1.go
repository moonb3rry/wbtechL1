package main

import (
	"fmt"
	"time"
)

// timer & select
func main() {
	ch := make(chan int)
	duration := 5 * time.Second // Длительность выполнения программы
	timer := time.NewTimer(duration)

	go func() {
		for i := 0; ; i++ {
			ch <- i // Отправка значений в канал
		}
	}()

	go func() {
		for {
			select {
			case <-timer.C: // Время истекло
				fmt.Println("Время истекло")
				close(ch) // Закрыть канал для остановки чтения
				return
			case value := <-ch: // Получение значения из канала
				fmt.Println(value)
			}
		}
	}()

	// Ожидание истечения таймера перед выходом из программы
	<-timer.C
}
