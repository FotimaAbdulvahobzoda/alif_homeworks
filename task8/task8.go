package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	resp1, err := client.Get("https://httpbin.org/delay/2") // Get "https://httpbin.org/delay/2": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	if err != nil {
		fmt.Println("Error!:", err)
	}
	defer resp1.Body.Close()
	fmt.Println("Status:", resp1.Status)

	client2 := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp2, err := client2.Get("https://httpbin.org/delay/2")
	if err != nil {
		fmt.Println("Error!:", err)
	}
	defer resp2.Body.Close()
	fmt.Println("Status:", resp2.Status) //Status: 200 OK
}
