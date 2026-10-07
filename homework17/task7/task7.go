package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	input := `{"name":"Алия","age":20,"city":"Душанбе"}`

	var p Person
	err := json.Unmarshal([]byte(input), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n%d\n", p.Name, p.Age) //city field is simply ignored, so Алия 20
}
