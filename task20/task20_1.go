package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Разделение строки на слова
	words := strings.Fields(s)
	// Переворачивание массива слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// Объединение слов обратно в строку
	return strings.Join(words, " ")
}

func main() {
	original := "snow dog sun"
	reversed := reverseWords(original)
	fmt.Println(reversed) // вывод: "sun dog snow"
}
