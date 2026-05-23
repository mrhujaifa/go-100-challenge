package main

/*
===========================================
  Go 100 Challenge
  Problem: #007
  Level: 🟡 Medium
===========================================

Problem: TODO CLI App

Topic:
- slice (add, remove by index)
- for loop
- switch for menu
- continuous user input loop

Industry Use:
Task Management Tools / CLI Productivity Apps /
Project Management Systems

Rules:
- A running menu will show these options:
    1. Add Task
    2. Delete Task
    3. List Tasks
    4. Exit
- Add Task    => append to slice
- Delete Task => remove by task number
- List Tasks  => print all tasks numbered
- Exit        => exit the program
- If no tasks => print "No tasks found!"
- If invalid delete number => "Invalid task number"

Example Run:
  === TODO App ===
  1. Add Task
  2. Delete Task
  3. List Tasks
  4. Exit
  Choose: 1
  Enter task: Buy groceries
  Task added!

  Choose: 1
  Enter task: Read Golang book
  Task added!

  Choose: 3
  1. Buy groceries
  2. Read Golang book

  Choose: 2
  Enter task number to delete: 1
  Task deleted!

  Choose: 3
  1. Read Golang book

  Choose: 4
  Goodbye!
===========================================
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func listTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}
	for index, task := range tasks {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}
func addTask(tasks []string) []string {
	fmt.Print("Enter task: ")
	input := readInput()
	if input == "" {
		fmt.Println("Task cannot be empty!")
		return tasks
	}
	tasks = append(tasks, input)
	fmt.Println("Task added successfully!")
	return tasks
}

func deleteTask(tasks []string, index int) []string {

	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("Task deleted!")
	return tasks
}

func main() {
	tasks := []string{}

	for {
		fmt.Println("\n=== TODO App ===")
		fmt.Println("1. Add Task")
		fmt.Println("2. Delete Task")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Choose: ")

		choice := readInput()

		switch choice {
		case "1":
			tasks = addTask(tasks)
		case "2":
			listTasks(tasks)
			fmt.Println("Enter task number to delete:")
			input := readInput()

			// conver string to num
			var num int
			fmt.Sscan(input, &num)

			tasks = deleteTask(tasks, num-1)

		case "3":
			listTasks(tasks)
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
