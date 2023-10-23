package main

import (
	"fmt"
)

func main() {
	// Исходный список строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество с помощью map
	set := make(map[string]struct{}) // 'struct{}' не занимает места
	for _, s := range strings {
		set[s] = struct{}{}
	}

	fmt.Println("Set:")
	for k := range set {
		fmt.Println(k)
	}
}
