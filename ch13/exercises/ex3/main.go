package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Tested using
// curl http://localhost:8080
func main() {

	http.HandleFunc("/", myHandler)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      http.DefaultServeMux,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

type jsonLog struct {
	IpAddr string    `json:"ip_address"`
	Time   time.Time `json:"time"`
}

func JsonLogging(r *http.Request) *jsonLog {
	data := jsonLog{
		IpAddr: r.RemoteAddr,
		Time:   time.Now(),
	}
	return &data
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// check the request URL path prefix
	path := r.URL.Path

	fmt.Println("Path:", path)
	if strings.HasPrefix(path, "/Form") {
		fmt.Println("Json")
		// return json response
		jsonData, err := json.MarshalIndent(JsonLogging(r), "", " ")
		if err != nil {
			panic(err)
		}

		// Set Content-Type header to JSON
		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonData)
	} else {
		fmt.Println("Text")
		// return text response
		out := time.Now().Format(time.RFC3339)
		w.Write([]byte(out))

	}
}
