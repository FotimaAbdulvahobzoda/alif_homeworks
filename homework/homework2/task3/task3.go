package main

import "fmt"

func main() {
	var (
		fueltype string
		liters   int
		cost     int
	)

	fmt.Print("Fuel type (92/95/98/disel): ")
	fmt.Scan(&fueltype)
	fmt.Print("How many liters: ")
	fmt.Scan(&liters)

	if liters <= 0 {
		fmt.Println("Error! Please enter a number.")
		return
	}

	switch fueltype {
	case "92":
		cost = 10
		total := cost * liters
		fmt.Printf("It costs %v by %v liters", total, liters)
	case "95":
		cost = 12
		total := cost * liters
		fmt.Printf("It costs %v by %v liters", total, liters)
	case "98":
		cost = 15
		total := cost * liters
		fmt.Printf("It costs %v by %v liters", total, liters)
	case "disel":
		cost = 11
		total := cost * liters
		fmt.Printf("It costs %v by %v liters", total, liters)
	default:
		fmt.Println("Invalid fuel type!")

	}
}
