package storage

import (
	"encoding/json"
	"os"
	"task-tracker/internal/models"
)

const filename = "tasks.json"

func ReadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []models.Task{}, nil
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func WriteTasks(tasks []models.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
