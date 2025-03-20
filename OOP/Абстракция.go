package main

import "fmt"

type Printer interface {
	Print()
}

type ConsolePrinter struct{}

func (cp ConsolePrinter) Print() {
	fmt.Println("Printing to console")
}

type FilePrinter struct{}

func (fp FilePrinter) Print() {
	fmt.Println("Printing to file")
}

func main() {
	var printer Printer
	printer = ConsolePrinter{}
	printer.Print() // Printing to console

	printer = FilePrinter{}
	printer.Print() // Printing to file
}

//Абстракция — это выделение основных характеристик объекта и игнорирование несущественных деталей. В Go абстракция достигается через интерфейсы.
//Интерфейс Printer абстрагирует процесс печати.
//Реализации ConsolePrinter и FilePrinter скрывают детали печати в консоль и файл.
