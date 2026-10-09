package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Company struct {
	Name string `json:"name"`
}

type User struct {
	Name    string  `json:"name"`
	Company Company `json:"company"`
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(
		"https://jsonplaceholder.typicode.com/users",
	)
	if err != nil {
		fmt.Println("Get Response Error:", err)
		return
	}
	defer resp.Body.Close()

	var users []User

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		fmt.Println("JSON Error:", err)
		return
	}

	for _, user := range users {
		fmt.Printf(
			"Name: %s | Company: %s\n",
			user.Name,
			user.Company.Name,
		)
	}
}
