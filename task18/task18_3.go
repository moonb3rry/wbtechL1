package main

import (
	"fmt"
)

func main() {
	counter := make(chan int, 1)
	counter <- 0

	done := make(chan bool)
	for i := 0; i < 1000; i++ {
		go func() {
			lastCount := <-counter
			lastCount++
			counter <- lastCount
			done <- true
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	finalCount := <-counter
	fmt.Printf("Final counter value: %d\n", finalCount)
}
