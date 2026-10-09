package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://httpbin.org/headers",
		nil,
	)
	if err != nil {
		fmt.Println("HTTP Error!:", err)
	}
	req.Header.Set("X-Student", "Fotima")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request Error!:", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading Error!:", err)
	}
	fmt.Println("Status:", resp.Status)
	fmt.Println(string(body))
}
