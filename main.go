package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/internal/commands"
	"task-tracker/internal/models"
)

// const filename = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		println(("Add a command to use task-tracker"))
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker add <description>")
			return
		}
		description := os.Args[2]

		commands.AddTask(description)

	case "list":
		if len(os.Args) == 2 {
			commands.ListTasks()
		} else {
			filter := os.Args[2]
			status := models.Status(filter)
			switch status {
			case models.StatusIncomplete, models.StatusInProgress, models.StatusComplete:
				commands.ListTasks(status)
			default:
				fmt.Fprintf(os.Stderr, "Error: invalid status '%s'\n", filter)
				os.Exit(1)
			}
		}

	case "update":
		if len(os.Args) < 4 {
			println("Usage: update 1 Description_here")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid ID: %s must be a number\n", os.Args[2])
			return
		}

		description := os.Args[3]
		updateTaskErr := commands.UpdateTask(id, description)
		if updateTaskErr != nil {
			println("Err", updateTaskErr)
			return
		}
		fmt.Printf("Task updated successfully ")

	case "delete":
		if len(os.Args) < 3 {
			println("Usage: delete task_id_here")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Printf("Invalid ID: %s must be a number\n", os.Args[2])
			return
		}

		deleteTaskErr := commands.DeleteTasks(id)

		if deleteTaskErr != nil {
			println("Erorr is", deleteTaskErr)
			return
		}

		fmt.Printf("Deleted tasks successfully ")

	case "status":
		if len(os.Args) < 4 {
			println("Usage: update id_here status_here")
			return
		}

		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Printf("Invalid ID: %s must be a number\n", os.Args[2])
			return
		}

		updateErr := commands.UpdateStatus(id, os.Args[3])
		if updateErr != nil {
			fmt.Print("err in updating", updateErr)
		}
	}

}
