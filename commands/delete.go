package commands

import (
	"fmt"

	utils "github.com/bnmwag/go-todo-app/utils"
)

func DeleteTask(index int) {
	index--

	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks to delete")
		return
	}

	if index < 0 || index >= int(len(tasks)) {
		fmt.Println("Invalid index:", index+1)
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	err = utils.WriteTasks(tasks)
	if err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Printf("Task %d deleted from the list.\n", index+1)
}
