package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Email struct {
	Adress string
}

type Phone struct {
	Number  int
	Balance int
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Sending email to %s\n", msg, e.Adress, "\f")
	return nil
}
func (p *Phone) Send(msg string) error {
	fmt.Printf("Sending phone to %s\n", msg, p.Number, "\f")
	return nil
}
func polimorfism(s Sender) {
	s.Send("polimorfism")
}

func Notify(s Sender) {

	if err := s.Send("Notify messege"); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	//email, ok := i.(Email) // это и есть утверждление типов
	switch s.(type) { //такая штука назвается type switch (это только для утверждения типов)
	case *Email: // тут утверждение типов
		fmt.Println("send email")
	case *Phone:
		fmt.Println("send phone")

	}
	phone, ok := s.(*Phone) // это утверждение типа, такая штука нам нудеа только в функциях по тиму этой , если при какой то структуре нам нужно работтаь с переменной этой структуры
	if ok {
		fmt.Println(phone.Balance)
	}
	fmt.Println("sucsess")
}

//// это для пустого интефеса
//func Notify(i interface{}) {
//	switch i.(type) { //такая штука назвается type switch (это только для утверждения типов)
//	case int:
//		fmt.Println("число не поддреживается")
//	}
//	s, ok := i.(Sender) //проверка на принадледность к интерфесу
//	if !ok {
//		fmt.Println("Error1")
//		return
//	}
//	err := s.Send("сообщение пустого интерфейса") // предалось ли мообщениее
//	if err != nil {
//		fmt.Printf("Error2")
//		return
//	}
//	fmt.Printf("sucsess")
//
//}

func main() {
	//это через интерфейсное значение
	// var sender Sender
	// sender = &Email{"dev@gmail.com"}
	// sender.Send("Hello")

	email := &Email{"dev@gmail.com"}           //мы не моожем передавать в интерфес структуру, только узказатель
	phone := &Phone{Number: 7777, Balance: 10} //мы не моожем передавать в интерфес структуру, только узказатель
	//slInterface := [5]int64{1, 2, 3, 4, 5}
	//Notify(slInterface)

	polimorfism(email)
	polimorfism(phone)
	Notify(email)
	Notify(phone)
	fmt.Println("")
}
