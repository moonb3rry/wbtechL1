package main

import "fmt"

// Интерфейс, который должен быть адаптирован
type Cat interface {
	Meow() string
}

// Класс, который нуждается в адаптации
type AlienCat struct{}

func (a *AlienCat) Purr() string {
	return "purr-purr"
}

// Адаптер: адаптирует вызовы к AlienCat, чтобы соответствовать интерфейсу Cat
type CatAdapter struct {
	alienCat *AlienCat
}

func (c *CatAdapter) Meow() string {
	return c.alienCat.Purr()
}

func main() {
	// Создание экземпляра AlienCat
	alienCat := &AlienCat{}

	// Адаптация AlienCat к интерфейсу Cat с использованием CatAdapter
	adapter := &CatAdapter{
		alienCat: alienCat,
	}

	// Теперь мы можем взаимодействовать с AlienCat, используя интерфейс Cat
	fmt.Println("AlienCat says:", adapter.Meow()) // AlienCat says: purr-purr
}
