package main

import (
	"fmt"
	"sync"
)

type command struct {
	key    int
	value  int
	action int // 1 - запись, 2 - чтение
	result chan<- int
}

func main() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	commands := make(chan command)

	go func() {
		for cmd := range commands { // Обработчик команд
			if cmd.action == 1 { // Запись
				m[cmd.key] = cmd.value
			} else if cmd.action == 2 { // Чтение
				cmd.result <- m[cmd.key]
			}
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cmd := command{key: i, value: i * 2, action: 1}
			commands <- cmd
		}(i)
	}

	wg.Wait()

	// Можно добавить сюда чтение данных из map через канал команд, если требуется
	close(commands) // Закрытие канала команд

	fmt.Println("map:", m)
}
