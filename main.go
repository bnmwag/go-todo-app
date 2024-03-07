package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bnmwag/go-todo-app/commands" // Replace with the actual path to your commands package
)

func getTask() string {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo add <task>")
		os.Exit(1)
	}
	return os.Args[2]
}

func getIndex() int {
	if len(os.Args) < 3 {
		fmt.Println("Usage: todo delete <index>")
		os.Exit(1)
	}
	index, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid index: only numbers are allowed")
		os.Exit(1)
	}
	return index
}

func main() {
	if len(os.Args) < 2 {
		commands.ListTasks()
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "list", "l":
		commands.ListTasks()
	case "add", "a":
		commands.AddTask(getTask())
	case "delete", "d":
		commands.DeleteTask(getIndex())
	case "check", "c":
		commands.CheckTask(getIndex())
	case "uncheck", "u":
		commands.UncheckTask(getIndex())
	case "help", "h":
		fmt.Println(usage())
	case "version", "v":
		fmt.Println(version())
	default:
		fmt.Println("Unknown command. Use 'todo help' for help.")
		os.Exit(1)
	}

	os.Exit(0)
}

func usage() string {
	// Define help message here
	return `
todo - A simple todo list CLI application

Usage: todo <command> [arguments]
    Running just 'todo' without arguments will also display all tasks.

Commands:
    list, l                 Display all tasks
    add, a      <task>      Add a new task to the list
                            - Example: todo add "Write a todo-cli app in GO!"
    delete, d   <index>     Delete a task from the list
                            - Example: todo delete 3
    check, c    <index>     Mark a task as complete
                            - Example: todo check 1
    uncheck, u  <index>     Mark a task as incomplete
                            - Example: todo uncheck 1  

    help, h                 Display help information
    version, v              Display version information
`
}

func version() string {
	// Return version information here
	return "1.0.0" // Replace with the actual version
}
