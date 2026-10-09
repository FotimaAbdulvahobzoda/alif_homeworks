package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	done := 0
	undone := 0
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Error!:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Server's response:", resp.Status)
	}

	var todo Todo
	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		fmt.Println("Decoding Error!:", err)
	}

	if todo.Completed {
		done += 1
	} else {
		undone += 1
	}

	fmt.Println("Completed:", done)
	fmt.Println("Not completed:", undone)

}
