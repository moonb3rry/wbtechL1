package main

import (
	"fmt"
	"sort"
)

// Для пользовательских типов данных, используя интерфейс sort.Interface
// Определение структуры Person
type Person struct {
	Name string
	Age  int
}

// Определение слайса типа Person, который будет реализовывать интерфейс sort.Interface
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println("До сортировки:", people)
	sort.Sort(ByAge(people)) // Сортировка по возрасту
	fmt.Println("После сортировки:", people)
}
