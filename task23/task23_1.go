package main

import "fmt"

func removeElement(slice []int, i int) []int {
	// Проверка корректности индекса
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	mySlice = removeElement(mySlice, indexToRemove)
	fmt.Println(mySlice) // [1 2 4 5]
}
