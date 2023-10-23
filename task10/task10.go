package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для хранения температур, сгруппированных по диапазону.
	groupedTemperatures := make(map[int][]float64)

	// Перебираем температуры и группируем их.
	for _, temp := range temperatures {
		// Определяем ключ для группировки каждого диапазона в 10 градусов.
		// Это делается путем деления температуры на 10, а затем округления до ближайшего целого числа.
		group := int(math.Floor(temp/10)) * 10

		// Добавляем температуру в соответствующую группу в карте.
		groupedTemperatures[group] = append(groupedTemperatures[group], temp)
	}

	// Выводим результат.
	for group, temps := range groupedTemperatures {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
