package main

import "fmt"

func removeElementWithCopy(slice []int, i int) []int {
	// Проверка корректности индекса
	if i < 0 || i >= len(slice) {
		return slice
	}
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	indexToRemove := 2
	mySlice = removeElementWithCopy(mySlice, indexToRemove)
	fmt.Println(mySlice) // [1 2 4 5]
}

/*Первый метод прост и сохраняет порядок элементов, но может быть неэффективен для больших срезов
из-за необходимости копировать элементы. Второй метод эффективен, но не сохраняет порядок.
Третий метод использует встроенную функцию copy, которая может быть более эффективной для больших срезов,
сохраняя при этом порядок элементов.*/
