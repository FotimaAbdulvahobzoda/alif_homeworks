package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for in := 1; in <= 8; in++ {
		time.Sleep(time.Duration(in) * 50 * time.Millisecond)
		fmt.Printf("Simple: %d is over\n", in) //slower
	}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
			fmt.Printf("Goroutine: %d is over\n", id)
		}(i)
	} //faster
	wg.Wait()
}
