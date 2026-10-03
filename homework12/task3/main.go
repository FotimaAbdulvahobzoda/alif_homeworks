package main

import (
	"fmt"
	validator "validEmail/validEmail"
)

type Person struct {
	Name  string
	Email string
	Age   int
}

func main() {
	person := []Person{
		{Name: "Alex", Email: "alex@example.com", Age: 20},
		{Name: "Bob", Email: "bobexample.com", Age: 25},
		{Name: "Charlie", Email: "charlie@example", Age: 30},
		{Name: "David", Email: "david@example.com", Age: 15},
	}

	for _, p := range person {
		validemail := validator.IsValidEmail(p.Email)
		adult := validator.IsAdult(p.Age)
		if !validemail {
			fmt.Printf("There's something wrong with your email: %s\n", p.Email)
		} else if !adult {
			fmt.Printf("You're not an adult!: %d\n", p.Age)
		} else {
			fmt.Println("User checkup completed successfully!")
		}
	}
}
