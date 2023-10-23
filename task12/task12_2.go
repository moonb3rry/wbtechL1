package main

import (
	"fmt"
)

// Функция для проверки наличия строки в срезе
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	// Исходный список строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество с помощью среза
	var set []string
	for _, s := range strings {
		if !contains(set, s) {
			set = append(set, s)
		}
	}

	fmt.Println("Set:", set)
}
