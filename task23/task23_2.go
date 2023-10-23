package main

import "fmt"

func removeElementWithoutOrder(slice []int, i int) []int {
	// Проверка корректности индекса
	if i < 0 || i >= len(slice) {
		return slice
	}
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	mySlice = removeElementWithoutOrder(mySlice, indexToRemove)
	fmt.Println(mySlice) // [1 2 5 4] (порядок изменен)
}
