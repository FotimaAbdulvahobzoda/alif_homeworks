package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	input := `{"title":"Go для начинающих","author":"Иван Петров","year":2023}`

	var b Book
	err := json.Unmarshal([]byte(input), &b)
	if err != nil {
		fmt.Println("Unwrapping the json went wrong: ", err)
	}
	fmt.Printf("%s\n%s\n%d\n", b.Title, b.Author, b.Year)

}
