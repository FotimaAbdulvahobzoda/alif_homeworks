package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)

	strings1 := []string{"A1", "A2", "A3", "A4", "A5"}
	strings2 := []string{"B1", "B2", "B3", "B4", "B5"}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for _, s := range strings1 {
			ch <- s
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()

		for _, s2 := range strings2 {
			ch <- s2
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	count := 0

	for s := range ch {
		fmt.Println(s)
		count++
	}

	fmt.Println("Number of strings:", count)
}
