package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	start := time.Now()
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	duration1 := time.Since(start)

	ch2 := make(chan int, 5)
	go func() {
		for i := 1; i < 6; i++ {
			time.Sleep(200 * time.Millisecond)
			<-ch2
		}
	}()

	start2 := time.Now()

	for i := 1; i <= 6; i++ {
		ch2 <- i
	}
	duration2 := time.Since(start2)
	fmt.Println("First channel (5 values):", duration1)
	fmt.Println("Second channel(6 values):", duration2)
}
