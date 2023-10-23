package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация переменных a и b значениями больше, чем 2^20
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(21), nil)

	// Для наглядности увеличим значения
	a.Add(a, big.NewInt(10)) // a = 2^20 + 10
	b.Add(b, big.NewInt(20)) // b = 2^21 + 20

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Выполнение арифметических операций
	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	quotient := new(big.Int).Div(a, b)

	// Вывод результатов
	fmt.Println("Sum (a + b):", sum)
	fmt.Println("Difference (a - b):", diff)
	fmt.Println("Product (a * b):", prod)
	fmt.Println("Quotient (a / b):", quotient)
}
