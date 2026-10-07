package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	p := Person{
		Name: "Maria",
		Age:  34,
		Address: Address{
			City:   "London",
			Street: "1/2 Boston Hill",
		},
	}

	data, _ := json.MarshalIndent(p, "", " ")
	fmt.Println("----- Marshall Indent -----")
	fmt.Println(string(data))

	input := `{
	"name": "Maria",
	"age": 34,
	"address": {
	"city": "London",
	"street": "1/2 Boston Hill"
 }
}`
	var p2 Person
	json.Unmarshal([]byte(input), &p2)
	fmt.Println("----- Unmarshal -----")
	fmt.Println(p2.Address.City)

}
