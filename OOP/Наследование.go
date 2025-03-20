package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
	Animal // Встраивание структуры Animal
	Breed  string
}

func main() {
	dog := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "German Shepherd",
	}
	dog.Speak() // Rex makes a sound
}

//Go не поддерживает классическое наследование, но использует композицию для достижения похожих целей.
//Структура Dog "наследует" методы и поля структуры Animal через встраивание (композицию).
