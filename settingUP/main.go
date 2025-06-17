package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string

func addTask(task string) {
	tasks = append(tasks, task)
	fmt.Println("Задача добавлена!")
}

func listTasks() {
	fmt.Println("Список задач:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1 - Добавить задачу\n2 - Показать задачи\n3 - Выход")
		fmt.Print("Выберите действие: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Введите задачу: ")
			scanner.Scan()
			task := scanner.Text()
			addTask(task)
		case "2":
			listTasks()
		case "3":
			fmt.Println("Выход.")
			return
		default:
			fmt.Println("Неверный выбор.")
		}
	}
}
