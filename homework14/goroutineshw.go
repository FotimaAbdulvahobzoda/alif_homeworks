package main

import (
	"fmt"
	"time"
)

func printSlowly(name string) {
	fmt.Printf("Start!: %s\n ", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished!: %s\n", name)
}

func main() {
	start := time.Now()
	printSlowly("hello world!")
	printSlowly("Seven minutes of heaven")
	printSlowly("D<E<A<T<H<M<E<T<A<L")

	fmt.Printf("Total took: %v\n", time.Since(start))

}
