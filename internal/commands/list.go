package commands

import (
	"fmt"
	"task-tracker/internal/models"
	"task-tracker/internal/storage"
)

func ListTasks(filter ...models.Status) error {
	fmt.Println(filter, filter)
	tasks, err := storage.ReadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	if len(filter) > 0 {
		filtered := []models.Task{}
		for _, task := range tasks {
			if task.Status == string(filter[0]) {
				filtered = append(filtered, task)
			}
		}
		tasks = filtered
	}

	for _, task := range tasks {
		fmt.Printf("Id %d | Status: %s | Description: %s\n", task.ID, task.Status, task.Description)
	}
	return nil
}
