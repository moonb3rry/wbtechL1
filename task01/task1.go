package main

import (
	"fmt"
)

// Определение структуры Human с методом Speak
type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.Name, h.Age)
}

// Определение структуры Action, встраивающей структуру Human
type Action struct {
	Human      // встраивание структуры Human
	ActionName string
}

// Определение метода для структуры Action
func (a *Action) DoAction() {
	fmt.Printf("%s делает действие: %s.\n", a.Name, a.ActionName)
}

func main() {
	// Создание экземпляра структуры Action
	person := Action{
		Human: Human{ // инициализация встроенной структуры Human
			Name: "Алексей",
			Age:  30,
		},
		ActionName: "программирование",
	}

	person.Speak()    // вызов метода Speak из встроенной структуры Human
	person.DoAction() // вызов метода DoAction из структуры Action
}
