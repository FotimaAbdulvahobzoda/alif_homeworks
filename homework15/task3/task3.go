package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 10
	fmt.Println(<-ch) // all goroutines are asleep - deadlock!
	// Go itself tells that it'll be stuck forever so it stops
}
