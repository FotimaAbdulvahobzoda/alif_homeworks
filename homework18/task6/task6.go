package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	Userid int
	Id     int
	Title  string
	Body   string
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts?userId=3")
	if err != nil {
		fmt.Println("HTTP Error!:", err)
	}

	defer resp.Body.Close()

	var posts []Post

	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		fmt.Println("Decoding Error!:", err)
	}

	fmt.Println("Posts:", len(posts))

	for _, post := range posts {
		fmt.Println(post)
	}
}
