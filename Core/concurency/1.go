package main

import (
	"fmt"
	"time"
)

// addTask кладёт новую задачу в канал
func addTask(mailbox chan<- string, taskName string) {
	fmt.Printf("Кладём задачу в ящик: %s\n", taskName)
	mailbox <- taskName
}

// takeTask берёт задачу из канала и имитирует её выполнение
func takeTask(mailbox <-chan string) {
	// Считываем задачи из канала в цикле
	// (цикл завершится, когда канал будет закрыт и все данные прочтены)
	for task := range mailbox {
		fmt.Printf("Получена задача: %s. Начинаем выполнение...\n", task)
		time.Sleep(time.Second) // Имитация выполнения задачи
		fmt.Printf("Задача '%s' выполнена!\n", task)
	}
}

func main() {
	// Создаём канал для строк — «очередь задач»
	mailbox := make(chan string)

	// Запускаем горутину, которая будет брать задачи из канала
	go takeTask(mailbox)

	// Кладём несколько задач. Каждая «кладка задачи» — тоже горутина
	go addTask(mailbox, "Сверстать лендинг")
	go addTask(mailbox, "Написать тесты")
	go addTask(mailbox, "Оптимизировать базу данных")

	// Ждём немного, чтобы все успели положить задачи и взять их из канала
	time.Sleep(4 * time.Second)

	// Закрываем канал, чтобы горутина takeTask могла завершиться
	close(mailbox)

	fmt.Println("Офис может закрываться...")
	time.Sleep(time.Second) // Немного ждём, чтобы увидеть вывод
}
