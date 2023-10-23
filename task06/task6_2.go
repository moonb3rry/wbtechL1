package main

import (
	"context"
	"fmt"
	"time"
)

// использование контекста для управления жизненным циклом
func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // ожидание сигнала на остановку
			fmt.Println("Остановка по сигналу:", ctx.Err())
			return
		default:
			fmt.Println("Работник работает")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx)

	time.Sleep(5 * time.Second) // имитация работы
	cancel()                    // инициирует остановку
	time.Sleep(1 * time.Second) // дает время на "очистку"
}
