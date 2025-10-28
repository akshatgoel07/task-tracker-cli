package commands

import (
	"task-tracker/internal/models"
	"task-tracker/internal/storage"
)

func DeleteTasks(id int) error {
	tasks, err := storage.ReadTasks()

	if err != nil {
		println("Err in getting tasks")
		return err
	}

	found := false
	newTasks := []models.Task{}

	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}

		newTasks = append(newTasks, task)

	}
	if !found {
		println("No task with such id found")
		return nil
	}

	return storage.WriteTasks(newTasks)

}
