package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type SimplePrinter struct{}

func (sp SimplePrinter) Print() {
	fmt.Println("Printing...")
}

type SimpleScanner struct{}

func (ss SimpleScanner) Scan() {
	fmt.Println("Scanning...")
}

func main() {
	printer := SimplePrinter{}
	scanner := SimpleScanner{}

	printer.Print() // Printing...
	scanner.Scan()  // Scanning...
}

//Принцип разделения интерфейса (Interface Segregation Principle, ISP)
//Определение: Клиенты не должны зависеть от интерфейсов, которые они не используют. Интерфейсы должны быть небольшими и специфичными.
//Интерфейсы Printer и Scanner разделены, чтобы клиенты могли использовать только нужные функции.
