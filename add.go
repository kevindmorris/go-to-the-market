package main

func add(tasks *[]Task, title string) {
	id := len(*tasks) + 1
	*tasks = append(*tasks, Task{ID: id, Title: title, Complete: false})
}
