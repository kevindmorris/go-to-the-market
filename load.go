package main

import (
	"encoding/json"
	"os"
)

func load() []Task {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks
}
