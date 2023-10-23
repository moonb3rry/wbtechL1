package main

import (
	"fmt"
)

// функция partition разделяет слайс на две части по опорному элементу и возвращает индекс этого элемента
func partition(arr []int, low, high int) int {
	pivot := arr[high] // выбор опорного элемента
	i := low - 1       // индекс меньшего элемента

	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// функция quickSort реализует алгоритм быстрой сортировки
func quickSort(arr []int, low, high int) {
	if low < high {
		// pi - индекс опорного элемента, arr[pi] находится на правильном месте
		pi := partition(arr, low, high)

		// Рекурсивно сортируем элементы перед разделением и после разделения
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)
	fmt.Println("Исходный массив:", arr)

	quickSort(arr, 0, n-1)
	fmt.Println("Отсортированный массив:", arr)
}
