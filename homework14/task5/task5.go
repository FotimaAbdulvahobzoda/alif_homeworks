package main

import (
	"fmt"
	"sync"
	"time"
)

func countTo(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	countTo(10, &wg)
	fmt.Println("AGAIN!")
	wg.Add(1)
	countTo(34, &wg)
	fmt.Println("ONE MORE TIME!")
	wg.Add(1)
	countTo(25, &wg)

	wg.Wait()
	fmt.Println("FINISHED!")
}
