package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func CalculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 5},
	}
	fmt.Println("Total area:", CalculateTotalArea(shapes)) // Total area: 128.5
}

//Принцип открытости/закрытости (Open/Closed Principle, OCP)
//Определение: Программные сущности (классы, модули, функции) должны быть открыты для расширения, но закрыты для модификации.
//Интерфейс Shape позволяет добавлять новые фигуры (например, треугольник), не изменяя функцию CalculateTotalArea.
