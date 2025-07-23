package main

import (
	"fmt"
)

func ls(tasks []Task) {
	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = "x"
		}
		fmt.Printf("[%s] %s\n", status, task.Title)
	}
}
