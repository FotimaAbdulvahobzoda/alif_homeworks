package main

import "fmt"

func something() {
	var (
		name string
		age  int
	)
	fmt.Print("What's your name?")
	fmt.Scan(&name)

	fmt.Print("what's your age?")
	fmt.Scan(&age)

	fmt.Printf("%s, you're %d\n", name, age)
}
