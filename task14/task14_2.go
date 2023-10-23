package main

import (
	"fmt"
)

func main() {
	variables := []interface{}{123, "string", true, make(chan bool)}

	for _, variable := range variables {
		if intValue, ok := variable.(int); ok {
			fmt.Printf("Тип переменной: int, значение: %d\n", intValue)
		} else if stringValue, ok := variable.(string); ok {
			fmt.Printf("Тип переменной: string, значение: %s\n", stringValue)
		} else if boolValue, ok := variable.(bool); ok {
			fmt.Printf("Тип переменной: bool, значение: %t\n", boolValue)
		} else if _, ok := variable.(chan bool); ok {
			fmt.Printf("Тип переменной: chan bool\n")
		} else {
			fmt.Printf("Неизвестный тип\n")
		}
	}
}
