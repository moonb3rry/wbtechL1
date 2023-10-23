package main

import (
	"fmt"
	"math"
)

// Point представляет структуру для точки в двумерном пространстве.
type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo вычисляет расстояние от текущей точки до другой точки.
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Hypot(p.x-other.x, p.y-other.y)
}

// SetX устанавливает значение координаты x для точки.
func (p *Point) SetX(x float64) {
	p.x = x
}

// SetY устанавливает значение координаты y для точки.
func (p *Point) SetY(y float64) {
	p.y = y
}

// GetX возвращает значение координаты x точки.
func (p *Point) GetX() float64 {
	return p.x
}

// GetY возвращает значение координаты y точки.
func (p *Point) GetY() float64 {
	return p.y
}

func main() {
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	fmt.Printf("Distance between points: %f\n", point1.DistanceTo(point2))
}
