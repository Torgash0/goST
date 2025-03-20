package main

import "fmt"

type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Sparrow is flying")
}

type Ostrich struct{}

func (o Ostrich) Fly() {
	fmt.Println("Ostrich cannot fly")
}

func MakeBirdFly(bird Bird) {
	bird.Fly()
}

func main() {
	sparrow := Sparrow{}
	ostrich := Ostrich{}

	MakeBirdFly(sparrow) // Sparrow is flying
	MakeBirdFly(ostrich) // Ostrich cannot fly
}

//Определение: Объекты в программе должны быть заменяемы экземплярами их подтипов без изменения правильности работы программы.
//Sparrow и Ostrich реализуют интерфейс Bird.
//Функция MakeBirdFly работает с любым типом, реализующим Bird.
