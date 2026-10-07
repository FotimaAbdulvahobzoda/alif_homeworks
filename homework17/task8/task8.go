package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func main() {
	input := `{"name":"Ноутбук","price":4500.5}`

	var i Item
	json.Unmarshal([]byte(input), &i)
	fmt.Printf("%+v/n", i) //Count field is set to default - 0
}
