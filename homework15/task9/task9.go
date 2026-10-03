package main

import "fmt"

func generateNumbers(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
	close(ch) //panic: close of closed channel
}

func generateNumbers1(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
	ch <- 10 //panic: send on closed channel
}

func main() {
	ch := make(chan int)
	fmt.Println("Starting!")
	go generateNumbers(9, ch)

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Finished!")

	fmt.Println("Starting! #2")
	go generateNumbers1(9, ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Finished! #2")
}
