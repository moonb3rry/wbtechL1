package main

import (
	"bytes"
	"fmt"
)

func reverseString2(s string) string {
	var buffer bytes.Buffer
	// Обход строки в обратном порядке и запись в буфер
	for i := len(s) - 1; i >= 0; i-- {
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}

func main() {
	original := "главрыба — абырвалг" // Этот метод может работать некорректно с Unicode
	reversed := reverseString2(original)
	fmt.Println(reversed) // вывод может быть некорректным
}
