package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Print the HTTP status code
	fmt.Println("Status Code:", res.Status)

	// Print the content type
	fmt.Println("Content-Type:", res.Header.Get("Content-Type"))

	// Print the body of the response
	fmt.Println("Response Body:", string(body))
}
