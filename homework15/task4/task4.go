package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Have puted 3 int values to the channel")
	fmt.Println("Length of the channel:", len(ch))
	fmt.Println("Capacity of the channel:", cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("Length of the channel:", len(ch))   //1
	fmt.Println("Capacity of the channel:", cap(ch)) //3

}
