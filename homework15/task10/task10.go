package main

import "fmt"

func sendToStudents(students []string) <-chan string {
	ch := make(chan string)

	go func() {
		for _, name := range students {
			ch <- name
		}

		close(ch)
	}()

	return ch
}

func main() {
	names := []string{"Ashley", "Kate", "Hannah", "Michael", "Anthony"}

	result := []string{}

	ch := sendToStudents(names)

	for n := range ch {
		result = append(result, n)
	}

	fmt.Println("Names:", len(names))
	fmt.Println("Total:", result)
}
