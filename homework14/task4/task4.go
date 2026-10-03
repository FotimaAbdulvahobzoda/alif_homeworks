package main

import (
	"fmt"
	"sync"
	"time"
)

func greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Hello, %s!\n", name)
	time.Sleep(300 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	names := []string{"Alex", "Siri", "Gojo Satoru", "Doughael", "Tom", "Angela"}
	for _, name := range names {
		wg.Add(1)
		greet(name, &wg)
	}
	wg.Wait()

	fmt.Println("Finished!")
}
