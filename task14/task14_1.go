package main

import (
	"fmt"
)

func main() {
	variables := []interface{}{123, "string", true, make(chan bool)}

	for _, variable := range variables {
		switch v := variable.(type) {
		case int:
			fmt.Printf("Тип переменной: int, значение: %d\n", v)
		case string:
			fmt.Printf("Тип переменной: string, значение: %s\n", v)
		case bool:
			fmt.Printf("Тип переменной: bool, значение: %t\n", v)
		case chan bool:
			fmt.Printf("Тип переменной: chan bool\n")
		default:
			fmt.Printf("Неизвестный тип\n")
		}
	}
}
