package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	messages := make(chan string)

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		timeout := time.NewTimer(2 * time.Second)
		defer timeout.Stop()

		for {
			select {
			case <-ticker.C:
				messages <- "ping"

			case <-timeout.C:
				close(done)
				return
			}
		}
	}()

	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)

		case <-done:
			fmt.Println("Конец")
			return
		}
	}
}
