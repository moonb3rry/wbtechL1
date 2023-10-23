package main

import (
	"fmt"
)

func main() {
	set1 := []int{1, 3, 2, 4}
	set2 := []int{6, 4, 3, 7}

	intersectionMap := make(map[int]bool)
	for _, element := range set1 {
		intersectionMap[element] = true
	}

	var intersectionSet []int
	for _, element := range set2 {
		if _, found := intersectionMap[element]; found {
			intersectionSet = append(intersectionSet, element)
			// Чтобы предотвратить дублирование, если исходные множества могут содержать повторяющиеся элементы.
			delete(intersectionMap, element)
		}
	}

	fmt.Println("Intersection is: ", intersectionSet)
}
