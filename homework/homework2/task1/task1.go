package main

import "fmt"

func main1() {
	var (
		cardbalance int
		sum         int
	)
	fmt.Print("Enter your balance card: ")
	fmt.Scan(&cardbalance)
	fmt.Print("Enter the sum you want to get: ")
	fmt.Scan(&sum)

	if cardbalance <= 0 {
		fmt.Println("Balance cannot be negative")
		return
	}
	if sum <= 0 {
		fmt.Println("Error! Please enter the number")
		return

	} else if sum > cardbalance {
		fmt.Println("Not enough sources")
		return
	}

	total := cardbalance - sum

	fmt.Printf("Left: %d\n", total)
}
