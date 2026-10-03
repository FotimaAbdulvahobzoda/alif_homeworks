package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i <= 100_000; i++ {
		wg.Add(1)
		go func(in int) {
			wg.Done()
			fmt.Println(in)
		}(i)
	}
	wg.Wait()
	fmt.Println("Time:", time.Since(start))
}
