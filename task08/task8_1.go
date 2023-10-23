package main

import (
	"fmt"
)

func setBit(number int64, position uint, state bool) int64 {
	var mask int64
	if state {
		// Установить i-й бит в 1
		mask = 1 << position
		return number | mask
	}
	// Установить i-й бит в 0
	mask = ^(1 << position)
	return number & mask
}

func main() {
	var number int64 = 0 // Начальное значение; все биты равны 0

	// Пример установки битов
	number = setBit(number, 1, true) // Установить 1-й бит в 1
	fmt.Printf("%064b\n", number)

	number = setBit(number, 5, true) // Установить 5-й бит в 1
	fmt.Printf("%064b\n", number)

	number = setBit(number, 1, false) // Установить 1-й бит в 0
	fmt.Printf("%064b\n", number)

	number = setBit(number, 3, true) // Установить 3-й бит в 1
	fmt.Printf("%064b\n", number)

}
