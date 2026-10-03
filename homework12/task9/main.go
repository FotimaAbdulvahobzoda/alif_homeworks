package main

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID, Name string
}

func NewUser(name string) User {
	return User{
		ID:   uuid.NewString(),
		Name: name,
	}
}

func main() {
	users := []string{"Alex", "Max", "Helen", "Angela", "Steven"}

	for _, u := range users {
		fmt.Println(NewUser(u))
	}
}
