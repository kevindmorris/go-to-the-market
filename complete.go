package main

import (
	"fmt"
)

func complete(tasks *[]Task, id int) {
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Complete = true
			return
		}
	}
	fmt.Println("Task not found.")
}
