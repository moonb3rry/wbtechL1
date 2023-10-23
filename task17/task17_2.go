package main

import "fmt"

func binarySearch(data []int, target int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + (high-low)/2 // чтобы избежать возможного переполнения

		if data[mid] == target {
			return mid
		}
		if data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // если элемент не найден
}

func main() {
	data := []int{10, 20, 30, 40, 50}
	target := 40

	index := binarySearch(data, target)

	if index != -1 {
		fmt.Printf("Найдено значение %d на позиции индекса %d.\n", target, index)
	} else {
		fmt.Println("Значение не найдено.")
	}
}
