package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	v, ok := <-ch
	fmt.Println("Number:", v, "Does it exist in the channel?:", ok) // 1 true
	v2, ok2 := <-ch
	fmt.Println("Number:", v2, "Does it exist in the channel?:", ok2) // 2 true
	v3, ok3 := <-ch
	fmt.Println("Number:", v3, "Does it exist in the channel?:", ok3) // 3 true
	v4, ok4 := <-ch
	fmt.Println("Number:", v4, "Does it exist in the channel?:", ok4) // 0 false

}
