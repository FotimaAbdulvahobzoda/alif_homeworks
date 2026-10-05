package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Worker", id, "finished")
}

func main() {
	var wg sync.WaitGroup
	workers := []int{23, 45, 67, 13, 40}
	for _, w := range workers {
		wg.Add(1)
		go func() {
			worker(w, &wg)
		}()
	}

	wg.Wait()
}
