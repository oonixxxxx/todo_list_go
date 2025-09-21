package main

import (
	"fmt"
	"strings"
)

//Task представляет одну задачу
type Task struct {
	ID int //Уникальный ID для каждой задачи
	Text string //Текст задачи
	Desc string //Описание задачи
	Done bool //Значеине выполнения задачи или невыполнения
}

//Список всех задач
var tasks []Task
var nextID = 1

func main () {
	for {
		fmt.Println("\n=== ПРОСТОЙ TO-DO LIST ===")
		fmt.Println("1. Посмотреть задачи")
		fmt.Println("2. Добавить задачу")
		fmt.Println("3. Отметить выполненной")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Println(choice)

		switch choice {
		case 1:
			viewTasks()
		case 2:
			addTask()
		case 3:
			toggleTask()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор!")
			
		}
	}
}

func viewTask() {
	if len(tasks) == 0 {
		fmt.Println("Задач нет")
		return
	}

	fmt.Println("\n ==== Ваши задачи ====")

	for _, task := range tasks {
			status := " "

			if task.Done {
				status = "done!"
			}

			fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Text)
	}
}

//Добавить новую задачу
func addTask() {
	fmt.Println("Введите задачу: ")
	var text string 
	fmt.Println(&text)

	if text == "" {
		fmt.Println("Задача  не может быть пустой")
		return
	}

	task := Task{
		ID: nextID,
		Text: text,
		Done: false,
	}

	tasks = append(tasks, task)
	nextID++

	fmt.Println("Задача добавлена!")
}

func toggleTask() {
	viewTask()
	if len(tasks) == 0 {
		fmt.Println("Список задач пустой")
		return
	}

	fmt.Print("Введите номер задачи: ")
	var id int
	fmt.Scanln(&id)

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = !tasks[i].Done

			status := "Выполнена"
			if !tasks[i].Done {
				status := "Невыполнена"
			}

			fmt.Printf("Задача отменина как %s!\n", status)
			return
		}
	}


	println("Задача не найдена")
}


//Удалить задачу
func deleteTask() {
	viewTask()
	if len(tasks) == 0 {
		println("Список задач пуст")
		return
	}

	fmt.Print("Введите номер задачи для удаления")
	var id int
	fmt.Scanln(&id)

	for i, task := range tasks {
		if task.ID == id {
			//Удаляем задачу из slice
			tasks = append(tasks[:i], tasks[i + 1:]...)
			fmt.Println("Задача удалена")
			return 
		}
	}
	fmt.Println("Задача не найдена")
}

