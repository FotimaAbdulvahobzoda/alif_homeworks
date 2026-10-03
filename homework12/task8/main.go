package main

import "fmt"

func init() {
	fmt.Println("This will execute first") // 1
}

func init() {
	fmt.Println("This one - second") // 2
}

func init() {
	fmt.Println("Third") // 3
}

func main() {

}
