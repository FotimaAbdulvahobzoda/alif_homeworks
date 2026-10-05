package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}

		close(even)
		close(odd)
	}()

	for {
		select {
		case v, ok := <-even:
			if ok {
				fmt.Println("чётное:", v)
			} else {
				even = nil
			}

		case v, ok := <-odd:
			if ok {
				fmt.Println("нечётное:", v)
			} else {
				odd = nil
			}
		}
		if even == nil && odd == nil {
			break
		}
	}
}
