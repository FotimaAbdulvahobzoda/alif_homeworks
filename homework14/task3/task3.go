package main

import (
	"fmt"
	"sync"
	"time"
)

func printSlowly(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("Start!: %s\n ", name)
	fmt.Printf("Finished!: %s\n", name)
}
func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	printSlowly("At least something...", &wg)
	wg.Add(1)
	printSlowly("Come on please", &wg)
	wg.Add(1)
	printSlowly("Good", &wg)

	wg.Wait()
	fmt.Println("Total:", time.Since(start))
}
