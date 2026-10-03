package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Length of the channel:", len(ch))
	fmt.Println("Capacity of the channel:", cap(ch))

	ch <- 4 //all goroutines are asleep - deadlock!
}
