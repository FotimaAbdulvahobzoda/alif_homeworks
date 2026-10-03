package main

import "fmt"

func init() {
	fmt.Println("This will execute first")
}

func init() {
	fmt.Println("This one - second")
}

func init() {
	fmt.Println("Third")
}
