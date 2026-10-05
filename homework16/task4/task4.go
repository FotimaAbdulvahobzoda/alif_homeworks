package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case c := <-ch:
		fmt.Println("The data of the channel:", c)
	default:
		fmt.Println("There's no data - the channel is empty")
	}
}
