package main

import (
	"fmt"
)

// setBitAlternative устанавливает бит в заданной позиции напрямую без маски.
func setBit2(number int64, position uint, state bool) int64 {
	if state {
		// Устанавливаем бит напрямую
		return number | (1 << position)
	}
	// Сбрасываем бит напрямую
	return number &^ (1 << position)
}

func main() {
	var number int64 = 0 // Начальное значение; все биты равны 0

	// Пример установки битов
	number = setBit2(number, 1, true) // Установить 1-й бит в 1
	fmt.Printf("%064b\n", number)

	number = setBit2(number, 5, true) // Установить 5-й бит в 1
	fmt.Printf("%064b\n", number)

	number = setBit2(number, 1, false) // Установить 1-й бит в 0
	fmt.Printf("%064b\n", number)

	number = setBit2(number, 3, true) // Установить 3-й бит в 1
	fmt.Printf("%064b\n", number)

}
