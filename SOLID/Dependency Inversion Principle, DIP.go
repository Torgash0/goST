package main

import "fmt"

type Database interface {
	Save(data string)
}

type MySQL struct{}

func (m MySQL) Save(data string) {
	fmt.Println("Saving to MySQL:", data)
}

type PostgreSQL struct{}

func (p PostgreSQL) Save(data string) {
	fmt.Println("Saving to PostgreSQL:", data)
}

type Apps struct {
	db Database
}

func (a Apps) Run(data string) {
	a.db.Save(data)
}

func main() {
	mysql := MySQL{}
	postgres := PostgreSQL{}

	app1 := Apps{db: mysql}
	app1.Run("Hello, MySQL!") // Saving to MySQL: Hello, MySQL!

	app2 := Apps{db: postgres}
	app2.Run("Hello, PostgreSQL!") // Saving to PostgreSQL: Hello, PostgreSQL!
}

//Принцип инверсии зависимостей (Dependency Inversion Principle, DIP)
//Определение: Модули верхнего уровня не должны зависеть от модулей нижнего уровня. Оба типа модулей должны зависеть от абстракций.
//App зависит от абстракции Database, а не от конкретных реализаций (MySQL, PostgreSQL)
