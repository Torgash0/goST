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

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 10, Height: 5.0}

	printArea(circle)    // Area: 78.5
	printArea(rectangle) // Area: 50
}

//Полиморфизм позволяет объектам разных типов обрабатываться как объекты одного типа. В Go это достигается через интерфейсы.
//Интерфейс Shape определяет метод Area.
//Структуры Circle и Rectangle реализуют интерфейс Shape.
//Функция printArea может работать с любым типом, реализующим интерфейс Shape.
