package main

import "fmt"

func main() {
	counter1 := 0
	counter2 := 0
	ch := make(chan int, 3)

	for i := 0; i <= 5; i++ {
		select {
		case ch <- i:
			counter1++
			fmt.Println("The number", i, "is added")
		default:
			fmt.Println("The buffer is full- skipping the number", i)
			counter2++
		}
	}
	fmt.Println("Numbers added:", counter1)
	fmt.Println("Numbers skipped:", counter2)
}
