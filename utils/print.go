package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/bnmwag/go-todo-app/types"
)

func generateRandomInt(min, max int64) int64 {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return min + random.Int63n(max-min+1)
}

func PrintEmptyMessage() {
	emptyMessages := []string{
		"All clear! Well done!",
		"Nothing to worry about!",
		"Your list is empty!",
		"Looks like you're free!",
		"Get yourself a cup of coffee!",
		"Take a break and relax!",
		"Enjoy some free time!",
		"Time to unwind!",
		"You're all caught up!",
	}

	fmt.Println("\nYour list is empty! 🎉")
	fmt.Printf(color.BlackString("%s\n\n"), emptyMessages[generateRandomInt(0, int64(len(emptyMessages)-1))])
}

func PrintTable(tasks types.Todos) {
	headerFmt := color.New(color.FgBlack).SprintfFunc()

	tbl := table.New("index", "Task", "State", "Added At", "Completed At")
	tbl.WithHeaderFormatter(headerFmt).WithPadding(3)

	doneCell := func(v interface{}) string {
		if v.(bool) {
			return color.GreenString("Yes")
		}
		return color.RedString("No")
	}

	completedAtCell := func(v time.Time) string {
		if v.IsZero() {
			return "N/A"
		}
		return v.Format("2006-01-02 15:04:05")
	}

	for idx, task := range tasks {
		tbl.AddRow(idx+1, task.Task, doneCell(task.Done), task.CreatedAt.Format("2006-01-02 15:04:05"), completedAtCell(task.CompletedAt))
	}

	fmt.Println()
	tbl.Print()
	fmt.Println()
}
