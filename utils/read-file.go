package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bnmwag/go-todo-app/constants"
	"github.com/bnmwag/go-todo-app/types"
)

// ReadFromFile reads data from a JSON file
func ReadFromFile() ([]byte, error) {
	data, err := os.ReadFile(constants.TodosJSON)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", constants.TodosJSON, err)
	}
	return data, nil
}

// ReadTasks reads and unmarshals tasks from a JSON file
func ReadTasks() (types.Todos, error) {
	data, err := ReadFromFile()
	if err != nil {
		return nil, err
	}

	var tasks types.Todos
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error decoding tasks from %s: %w", constants.TodosJSON, err)
	}

	return tasks, nil
}
