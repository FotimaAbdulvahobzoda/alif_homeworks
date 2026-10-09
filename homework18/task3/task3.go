package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/99999",
	}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Response Error:", err)
		}

		fmt.Println("Status code:", resp.StatusCode)
		if resp.StatusCode == 200 {
			var todo Todo

			if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
				fmt.Println("Decoding Error:", err)
				resp.Body.Close()
			}
			fmt.Println("Title:", todo.Title)
		} else if resp.StatusCode == 404 {
			fmt.Println("Title Not Found")
		}

		resp.Body.Close()
	}
}
