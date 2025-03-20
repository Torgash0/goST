package main

import "fmt"

type Engine struct {
	Power int
}

type Cars struct {
	Engine // Композиция: Car "имеет" Engine
	Brand  string
}

func main() {
	car := Cars{
		Engine: Engine{Power: 150},
		Brand:  "Toyota",
	}
	fmt.Println("Car brand:", car.Brand)
	fmt.Println("Engine power:", car.Power)
}

//Эти концепции описывают отношения между объектами. В Go они реализуются через встраивание структур и использование указателей.
//Car содержит Engine как часть себя (композиция).
