package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	InStock bool    `json:"in_stock"`
}

func main() {
	p := Product{
		ID:      22,
		Name:    "Cheese",
		Price:   20,
		InStock: true,
	}

	data, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
