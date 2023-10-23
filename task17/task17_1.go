package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{10, 20, 30, 40, 50}
	target := 40

	// Поиск индекса target в слайсе data.
	// Если target не найден, функция вернет len(data), т.е. индекс, который target бы занял, будучи вставленным.
	index := sort.Search(len(data), func(i int) bool { return data[i] >= target })

	// Проверка, найден ли искомый элемент.
	if index < len(data) && data[index] == target {
		fmt.Printf("Найдено значение %d на позиции индекса %d.\n", target, index)
	} else {
		fmt.Println("Значение не найдено.")
	}
}
