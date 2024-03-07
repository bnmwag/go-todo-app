package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bnmwag/go-todo-app/constants"
	"github.com/bnmwag/go-todo-app/types"
)

// WriteToFile writes data to a JSON file
func WriteToFile(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ") // Indented for readability
	if err != nil {
		return fmt.Errorf("error encoding data: %w", err)
	}

	err = os.WriteFile(constants.TodosJSON, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file %s: %w", constants.TodosJSON, err)
	}

	return nil
}

// WriteTasks writes a slice of tasks to a JSON file
func WriteTasks(tasks types.Todos) error {
	return WriteToFile(tasks)
}
