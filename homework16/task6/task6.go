package main

import "fmt"

func tryRead(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Println("Message:", msg)
	default:
		fmt.Println("Empty")
	}
}

func main() {
	ch := make(chan string, 1)

	ch <- "There's definetly something"

	tryRead(ch) //Message: There's definetly something
	tryRead(ch) //Empty
	tryRead(ch) //Empty

}
