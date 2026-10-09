package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Users struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("HTTP Error!:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Server's response:", resp.Status)
	}

	var user Users
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Decode error!:", err)
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Phone:", user.Phone)

}
