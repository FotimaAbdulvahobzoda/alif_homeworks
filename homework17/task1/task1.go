package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

func main() {
	b := Book{
		Title:  "Diary of Whimpy Kid",
		Author: "Jeff Kinney",
		Year:   2013,
		Price:  7.99,
	}
	data, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error!", err)
	}
	fmt.Println(string(data))
}
