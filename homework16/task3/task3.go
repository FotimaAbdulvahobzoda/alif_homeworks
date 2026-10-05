package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch1 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 42
	}()

	select {
	case msg := <-ch:
		fmt.Println("Have waited for:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("TIMEOUT!")
	}

	select {
	case msg := <-ch1:
		fmt.Println("Have waited for:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("TIMEOUT!")
	}
}
