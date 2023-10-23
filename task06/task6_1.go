package main

import (
	"fmt"
	"time"
)

// использование канала для сигнала остановки
func worker(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Остановка работника")
			return // останавливает горутину
		default:
			fmt.Println("Работник работает")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop := make(chan bool)
	go worker(stop)

	time.Sleep(5 * time.Second) // имитация работы
	stop <- true                // отправляем сигнал остановки
	time.Sleep(1 * time.Second) // даем время для "очистки"
}
