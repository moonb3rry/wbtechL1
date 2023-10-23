package main

import (
	"fmt"
	"time"
)

// использование канала для остановки по выполненной задаче
func worker3(done chan<- bool) {
	fmt.Println("Работник начинает работу")
	time.Sleep(5 * time.Second) // Представим, что работник что-то делает
	fmt.Println("Работник закончил работу")

	done <- true // сигнализируем, что работа завершена
}

func main() {
	done := make(chan bool, 1)
	go worker3(done)

	<-done // ожидание сигнала о завершении работы
}
