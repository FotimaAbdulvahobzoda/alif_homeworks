package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	u := User{
		Name:  "Hannah",
		Age:   23,
		Email: "HannahMontanna@gmail.com",
	}

	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error!:", err)
	}
	fmt.Println(string(data))
}
