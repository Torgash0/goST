package main

import "fmt"

type BankAccount struct {
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
	}
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= b.balance {
		b.balance -= amount
	}
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func main() {
	account := BankAccount{}
	account.Deposit(1000)
	account.Withdraw(500)
	fmt.Println("Balance:", account.Balance()) // Balance: 500
}

//Инкапсуляция — это сокрытие внутренней реализации объекта и предоставление доступа только через публичные методы.
//Поле balance инкапсулировано (не экспортируется, так как начинается со строчной буквы).
//Доступ к balance осуществляется через методы Deposit, Withdraw и Balance.
