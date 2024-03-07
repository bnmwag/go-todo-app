package commands

import (
	"fmt"
	"os"

	utils "github.com/bnmwag/go-todo-app/utils"
)

func ListTasks() {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		utils.PrintEmptyMessage()
		os.Exit(0)
	}

	utils.PrintTable(tasks)
}
