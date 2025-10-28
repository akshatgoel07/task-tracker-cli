package commands

import (
	"fmt"
	"task-tracker/internal/models"
	"task-tracker/internal/storage"
)

func UpdateStatus(id int, status string) error {

	if status != string(models.StatusIncomplete) && status != string(models.StatusComplete) && status != string(models.StatusInProgress) {
		return fmt.Errorf("invalid status: must be one of %s, %s, or %s",
			models.StatusIncomplete, models.StatusComplete, models.StatusInProgress)
	}

	if len(status) == 0 {
		println("Incorrect syntax for status update")
		return nil
	}

	tasks, err := storage.ReadTasks()

	if err != nil {
		return fmt.Errorf("failed to read tasks: %w", err)
	}

	found := false

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status

			found = true
			break
		}
	}
	if !found {
		println("Task id not found")
		return nil
	}

	writeErr := storage.WriteTasks(tasks)

	if writeErr != nil {
		return fmt.Errorf("failed to write tasks %w", writeErr)
	}

	return nil
}
