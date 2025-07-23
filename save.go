package main

import (
	"encoding/json"
	"log"
	"os"
)

func save(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		log.Fatal("Failed to encode tasks:", err)
		return
	}
	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		log.Printf("Error writing to file %s: %v\n", taskFile, err)
		return
	}
}
