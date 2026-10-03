package main

import "fmt"

func main() {
	var salary float64

	fmt.Print("Enter your salary: ")
	fmt.Scan(&salary)

	if salary <= 0 {
		fmt.Println("Error! Salary cannot be zero")
		return
	}

	if salary <= 3000 {
		order := salary * 0.5
		sumafterorder := salary - order
		fmt.Printf("Your tax: %v\nYour sum after tax: %v\n", order, sumafterorder)
	} else if salary >= 3000 && salary <= 10000 {
		order := salary * 0.10
		sumafterorder := salary - order
		fmt.Printf("Your tax: %v\nYour sum after tax: %v\n", order, sumafterorder)
	} else if salary >= 10000 {
		order := salary * 0.15
		sumafterorder := salary - order
		fmt.Printf("Your tax: %v\nYour sum after tax: %v\n", order, sumafterorder)
	}
}
