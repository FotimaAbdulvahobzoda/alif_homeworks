package main

import "fmt"

func main() {
	for i := 100; i >= 90; i-- {
		fmt.Println(i, "")
	}
	task7()
}

func task7() {
	num := 1234
	counter := 0

	if num == 0 {
		counter = 1
	} else {
		for num > 0 {
			counter++
			num /= 10
		}
	}

}
