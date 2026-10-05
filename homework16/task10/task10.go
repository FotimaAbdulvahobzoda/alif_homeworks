package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	results := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()

		time.Sleep(300 * time.Millisecond)
		results <- 1
	}()

	go func() {
		defer wg.Done()

		time.Sleep(500 * time.Millisecond)
		results <- 2
	}()

	go func() {
		defer wg.Done()

		time.Sleep(700 * time.Millisecond)
		results <- 3
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0

	for v := range results {
		fmt.Println("Received:", v)
		sum += v
	}

	fmt.Println("Sum:", sum)
}
