package main

import "fmt"

type Driver struct {
	Name string
}

type Car struct {
	Driver *Driver // Ассоциация: Car "связан" с Driver
	Brand  string
}

func main() {
	driver := &Driver{Name: "John"}
	car := Car{
		Driver: driver,
		Brand:  "Toyota",
	}
	fmt.Println("Car brand:", car.Brand)
	fmt.Println("Driver name:", car.Driver.Name)
}

//Car связан с Driver через указатель (ассоциация).
