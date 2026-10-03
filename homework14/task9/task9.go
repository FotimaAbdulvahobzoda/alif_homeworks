package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 50; i++ {
		wg.Add(1)
		fmt.Println(runtime.NumGoroutine())
		go func(s int) {
			wg.Done()
			fmt.Println(s)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
}
