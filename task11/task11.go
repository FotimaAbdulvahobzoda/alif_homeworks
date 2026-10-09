package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		"https://jsonplaceholder.typicode.com/posts/1",
		nil,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request Error!:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
}
