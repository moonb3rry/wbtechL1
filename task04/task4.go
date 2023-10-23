package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(1 * time.Second) // Имитация обработки данных
	}
}

func main() {
	// Инициализация канала для обработки сигналов прерывания (Ctrl+C)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Конфигурация воркеров
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scanln(&numWorkers)

	jobs := make(chan int, 100) // Буферизованный канал для заданий
	var wg sync.WaitGroup

	// Создание воркеров
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Горутина для непрерывной отправки данных в канал
	go func() {
		counter := 1
		for {
			select {
			case <-stopChan: // При получении сигнала прерывания
				fmt.Println("\nПолучен сигнал завершения. Остановка отправки работ.")
				close(jobs) // Закрытие канала для корректного завершения воркеров
				return
			default:
				jobs <- counter
				counter++
				time.Sleep(500 * time.Millisecond) // Имитация задержки при получении данных
			}
		}
	}()

	// Ожидание сигнала завершения (Ctrl+C)
	<-stopChan

	// Дождаться завершения всех воркеров
	wg.Wait()
	fmt.Println("Программа завершила работу.")
}
