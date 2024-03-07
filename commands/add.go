package commands

import (
	"fmt"
	"time"

	types "github.com/bnmwag/go-todo-app/types"
	utils "github.com/bnmwag/go-todo-app/utils"
)

func AddTask(task string) {
	tasks, err := utils.ReadTasks()
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	newTask := types.Todo{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	tasks = append(tasks, newTask)

	err = utils.WriteTasks(tasks)
	if err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Printf("Task '%s' added to the list.\n", task)
}
