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
	go printSlowly("hello world!")            //this
	go printSlowly("Seven minutes of heaven") //yes and no
	go printSlowly("D<E<A<T<H<M<E<T<A<L")     //this one either yes or no

	fmt.Printf("Total took: %v\n", time.Since(start)) //task 3 starts

	time.Sleep(2 * time.Second) //takes too long

}
