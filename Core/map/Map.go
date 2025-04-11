package main

import (
	"fmt"
)

func main() {
	fmt.Println("MAP")
	//MAP
	var m1 map[int32]int     //создание мапы
	var m2 map[string]string //создание мапы
	m3 := make(map[int]int)  //создание мапы
	ages := map[string]int{  //создание мапы
		"Andru": 20,
		"Mike":  19,
	}
	age := ages["Andru"] // так мы получим значение по ключу
	ages["Vera"] = 23    // добавление нового значения
	agesln := len(ages)
	age, ageExists := ages["Vera"] //проверка есть ли в мапе этот ключ
	fmt.Println(age, ageExists)
	age, ageExists = ages["0"]
	fmt.Println(m1, m2, m3, ages, age, ages, agesln)
	delete(ages, "Andre") //удаление элемента из мапы
}
