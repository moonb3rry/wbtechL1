package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Преобразование строки в срез рун, чтобы корректно обработать Unicode символы
	runes := []rune(s)
	// Переворачивание среза рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Возвращение результата как строки
	return string(runes)
}

func main() {
	original := "главрыба — абырвалг"
	reversed := reverseString(original)
	fmt.Println(reversed) // вывод: "главрываб — абывралг"
}
