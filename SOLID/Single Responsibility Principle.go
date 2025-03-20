package main

import (
	"fmt"
	"os"
)

// Отвечает только за работу с файлами
type FileManager struct{}

func (fm FileManager) SaveToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

// Отвечает только за логику приложения
type App struct {
	fileManager FileManager
}

func (a App) Run() {
	data := []byte("Hello, SOLID!")
	if err := a.fileManager.SaveToFile("output.txt", data); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data saved successfully")
	}
}

func main() {
	app := App{fileManager: FileManager{}}
	app.Run()
}

//Принцип единственной ответственности (Single Responsibility Principle, SRP)
//Определение: Класс (или структура в Go) должен иметь только одну причину для изменения, то есть выполнять только одну задачу.
//FileManager отвечает только за работу с файлами.
//App отвечает только за логику приложения.
//Каждая структура имеет одну ответственность.
