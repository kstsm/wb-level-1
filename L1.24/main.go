package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками на плоскости.
//Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
//Расстояние рассчитывается по формуле между координатами двух точек.
//Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.

// Point структура, представляющая точку на плоскости
type Point struct {
	x float64
	y float64
}

// NewPoint конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance метод для вычисления расстояния до другой точки
func (p Point) Distance(other Point) float64 {
	// Формула расстояния между двумя точками
	dx := p.x - other.x
	dy := p.y - other.y

	//Возвращаем квадратный корень из этого значения
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаём две точки через конструктор
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние
	fmt.Printf("Расстояние между точками: %.2f\n", p1.Distance(p2))
}
