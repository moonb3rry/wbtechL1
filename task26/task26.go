package main

import (
	"fmt"
	"strings"
)

// isUnique проверяет, все ли символы в строке уникальные (регистронезависимо).
func isUnique(s string) bool {
	// Приводим строку к нижнему регистру для регистронезависимости
	s = strings.ToLower(s)

	// Создаем карту для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	// Проверяем каждый символ в строке
	for _, char := range s {
		if _, exists := charMap[char]; exists {
			// Если символ уже встречался, возвращаем false
			return false
		}
		// Если символ уникален, добавляем его в карту
		charMap[char] = true
	}

	// Если дубликатов не найдено, возвращаем true
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))      // true
	fmt.Println(isUnique("abCdefAaf")) // false
	fmt.Println(isUnique("\taabcd"))   // false
}
