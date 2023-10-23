package main

import (
	"fmt"
	"sort"
)

// Для слайсов встроенных типов (например, []int, []float64, []string)
func main() {
	// Для слайса int
	ints := []int{44, 67, 3, 17, 89, 10, 73, 9, 21, 50}
	sort.Ints(ints)
	fmt.Println("Сортировка слайса ints:", ints)

	// Для слайса float64
	floats := []float64{12.7, 43.2, 1.9, 5.7, 30.1, 5.3, 40.8, 18.6}
	sort.Float64s(floats)
	fmt.Println("Сортировка слайса floats:", floats)

	// Для слайса string
	strings := []string{"banana", "apple", "cherry", "date", "fig", "grape"}
	sort.Strings(strings)
	fmt.Println("Сортировка слайса strings:", strings)
}
