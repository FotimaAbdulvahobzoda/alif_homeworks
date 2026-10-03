package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)
	go square(7, ch)
	fmt.Println("number 7 in power of 2:", <-ch)
}
