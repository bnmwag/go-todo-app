package commands

import (
	"fmt"
	"time"

	utils "github.com/bnmwag/go-todo-app/utils"
)

func CheckTask(index int) {
	index--

	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks to check")
		return
	}

	if index < 0 || index >= int(len(tasks)) {
		fmt.Println("Invalid index:", index+1)
		return
	}

	if tasks[index].Done {
		fmt.Printf("Task %d is already checked.\n", index+1)
		return
	}

	tasks[index].Done = true
	tasks[index].CompletedAt = time.Now()

	err = utils.WriteTasks(tasks)
	if err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Printf("Task %d checked.\n", index+1)
}
