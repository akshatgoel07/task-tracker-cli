# Task Tracker

A simple command-line task tracker built in Go.


## Usage

```bash
# Add a new task
./task-tracker add "Complete project documentation"

# List all tasks
./task-tracker list

# List tasks by status
./task-tracker list incomplete
./task-tracker list in-progress
./task-tracker list complete

# Update a task description
./task-tracker update 1 "Updated task description"

# Update task status
./task-tracker status 1 complete

# Delete a task
./task-tracker delete 1
```

## Status Values

- `Incomplete` - Task not started
- `In Progress` - Task in progress
- `Complete` - Task completed


