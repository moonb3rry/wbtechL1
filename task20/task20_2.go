package main

import (
	"fmt"
	"strings"
)

func reverseWords2(s string) string {
	words := strings.Fields(s)
	stack := make([]string, 0, len(words))

	// Помещение слов в стек
	for _, word := range words {
		stack = append(stack, word)
	}

	// Извлечение слов из стека в обратном порядке
	reversedWords := make([]string, 0, len(words))
	for len(stack) > 0 {
		// Извлечение верхнего элемента стека
		n := len(stack) - 1
		reversedWords = append(reversedWords, stack[n])
		stack = stack[:n]
	}

	// Объединение слов обратно в строку
	return strings.Join(reversedWords, " ")
}

func main() {
	original := "snow dog sun"
	reversed := reverseWords2(original)
	fmt.Println(reversed) // вывод: "sun dog snow"
}
