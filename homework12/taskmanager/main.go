package main

import (
	"fmt"
	"taskmanager/task"
)

func main() {
	task := []task.Task{
		{Title: "Do homework", Done: false},
		{Title: "Cook meal", Done: true},
		{Title: "Understanding Go", Done: false},
		{Title: "Hate chemistry", Done: true},
	}

	for _, t := range task {
		if t.Done {
			fmt.Println("[x] ", t.Title)
		} else {
			fmt.Println("[ ] ", t.Title)
		}
	}
}
