package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	codes := []int{200, 201, 404, 500, 503}
	for _, code := range codes {
		url := fmt.Sprintf("https://httpbin.org/status/%d", code)
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("HTTP Error!:", err)
			continue
		}
		switch {
		case resp.StatusCode >= 200 && resp.StatusCode < 300:
			fmt.Printf("%d: victory!\n", code)
		case resp.StatusCode >= 400 && resp.StatusCode < 500:
			fmt.Printf("%d: client error\n", code)
		case resp.StatusCode >= 500:
			fmt.Printf("%d: server error\n", code)
		}
		resp.Body.Close()
	}

}
