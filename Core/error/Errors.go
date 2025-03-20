package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNotFound = errors.New("not found")

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	//логирование
	if err != nil {
		log.Println("Ошибка", err)
	} else {
		fmt.Println(result)
	}

	//проверка типа ошибки
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Элемент не найден")
	} else {
		fmt.Println("Ошибок нет или другая ошибка")
	}
}
