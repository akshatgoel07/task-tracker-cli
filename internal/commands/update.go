package commands

import (
	"fmt"
	"task-tracker/internal/storage"
	"time"
)

func UpdateTask(id int, description string) error {

	if len(description) == 0 {
		return fmt.Errorf("not valid description")
	}
	// get all, filter all del existing, create new, return it
	tasks, err := storage.ReadTasks()

	if err != nil {
		return err
	}
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return storage.WriteTasks(tasks)

}
