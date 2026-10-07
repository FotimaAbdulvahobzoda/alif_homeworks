package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Comment string `json:"comment,omitempty"`
}

func main() {
	s1 := Student{
		Name:    "Alicia",
		Grade:   9,
		Comment: "Hi! Nice to meet you!",
	}

	s2 := Student{
		Name:  "Matthew",
		Grade: 8,
	}

	data1, _ := json.MarshalIndent(s1, "", " ")
	data2, _ := json.MarshalIndent(s2, "", " ")
	fmt.Println(string(data1))
	fmt.Println(string(data2)) //it doesn't have comment in json
}
