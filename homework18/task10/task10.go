package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type PostResponse struct {
	ID int `json:"id"`
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	post := Post{
		Title:  "Learning Go",
		Body:   "I'm so stupid please help",
		UserID: 1,
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		fmt.Println("JSON Error:", err)
		return
	}

	resp, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("Request Error:", err)
		return
	}
	defer resp.Body.Close()

	var result PostResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("JSON response Error:", err)
		return
	}

	fmt.Println("Statud code:", resp.StatusCode)
	fmt.Println("ID:", result.ID)
}
