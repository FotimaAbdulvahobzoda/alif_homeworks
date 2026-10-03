package main

import (
	"fmt"
	strutil "word/word"
)

func main() {
	words := []string{"level", "radar", "Golang", "update", "mom", "yogurt", "dad", "fan"}

	for _, w := range words {
		boo := strutil.IsPalidrome(w)
		if boo {
			fmt.Printf("\n%s is palidrome!", w)
		} else {
			fmt.Printf("\n%s is sadly not a palidrome :(", w)
		}
	}

}
