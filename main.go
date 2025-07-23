package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

const taskFile = "db.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <add|ls|complete> [task]")
		return
	}

	command := os.Args[1]
	tasks := load()

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		add(&tasks, title)
		save(tasks)
		fmt.Println("Task added.")
	case "ls":
		ls(tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task ID to mark as complete.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}
		complete(&tasks, id)
		save(tasks)
		fmt.Println("Task marked as complete.")
	default:
		fmt.Println("Unknown command:", command)
	}
}
