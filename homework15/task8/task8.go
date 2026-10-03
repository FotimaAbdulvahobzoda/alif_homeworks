package main

import "fmt"

func generateNumbers(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	fmt.Println("Starting!")
	go generateNumbers(9, ch)

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Finished!")
}
