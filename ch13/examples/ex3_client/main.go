package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func printHelpText() {
	fmt.Printf(`		
Usage:
	./program

Optional Args:
	--json | -j - if you want the value returned in json format
`)
}

func parseArgs(args []string) bool {
	if len(args) > 2 {
		printHelpText()
		os.Exit(0)
	}

	useJson := false
	if len(args) == 2 {
		switch args[1] {
		case "--json", "-j":
			useJson = true
		default:
			printHelpText()
		}

	}
	return useJson
}

func main() {

	useJson := parseArgs(os.Args)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	var ipAddress string = "http://localhost"
	var port string = "8080"
	ipAddress = fmt.Sprintf("%s:%s", ipAddress, port)

	fmt.Println("Ip", ipAddress)
	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, ipAddress, nil)
	if err != nil {
		panic(err)
	}

	if useJson {
		fmt.Println("Adding header")
		req.Header.Set("Accept", "application/json") // Set Accept header to request JSON response
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
