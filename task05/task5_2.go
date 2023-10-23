package main

import (
	"fmt"
	"time"
)

// time.after
func main() {
	ch := make(chan int)
	duration := 5 * time.Second

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	go func() {
		timeout := time.After(duration) // Автоматический таймер
		for {
			select {
			case <-timeout:
				fmt.Println("Время истекло")
				close(ch)
				return
			case value := <-ch:
				fmt.Println(value)
			}
		}
	}()

	<-time.After(duration) // Подождать, пока не истечет время
}
