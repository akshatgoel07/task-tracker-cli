package commands

import (
	"fmt"
	"task-tracker/internal/models"
	"task-tracker/internal/storage"
	"time"
)

func AddTask(description string) error {
	if len(description) == 0 {
		return fmt.Errorf("task description cannot be empty")
	}

	tasks, err := storage.ReadTasks()
	if err != nil {
		return err
	}

	newTask := models.Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      string(models.StatusIncomplete),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	fmt.Printf("Task added with id %d", newTask.ID)
	return storage.WriteTasks(tasks)
}
