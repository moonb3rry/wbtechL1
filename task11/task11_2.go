package main

import (
	"fmt"
)

func main() {
	set1 := []int{1, 3, 2, 4}
	set2 := []int{6, 4, 3, 7}

	var intersectionSet []int
	for _, element1 := range set1 {
		for _, element2 := range set2 {
			if element1 == element2 {
				intersectionSet = append(intersectionSet, element1)
				break // Переходим к следующему элементу set1, так как пересечение уже найдено.
			}
		}
	}

	fmt.Println("Intersection is: ", intersectionSet)
}
