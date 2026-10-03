package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine!"
	}()
	v := <-ch
	fmt.Println(v)
}
