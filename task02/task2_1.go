package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // Для синхронизации и ожидания завершения всех горутин

	for _, num := range numbers {
		wg.Add(1) // Указываем, что начинается новая горутина
		go func(n int) {
			defer wg.Done() // По завершении горутины сигнализируем об этом
			fmt.Println(n * n)
		}(num)
	}

	wg.Wait() // Ожидание завершения всех горутин
}

// Преимущества:
// - Простота и читаемость: этот подход легко читать и понимать, он идеально подходит для простых конкурентных задач.
// - Контроль над горутинами: вы знаете, когда каждая горутина начинает и заканчивает работу.
//
// Недостатки:
// - Менее эффективен для больших объемов задач: создание горутины для каждой маленькой задачи может быть ресурсоемким.
// - Отсутствие встроенной очереди задач: необходимо самостоятельно управлять очередностью и балансом задач.