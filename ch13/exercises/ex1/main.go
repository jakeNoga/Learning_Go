package main

import (
	"net/http"
	"time"
)

// ex1:
// Write a small web server that returns the current time in RFC 3339 format when you send it a GET command. You can use a third-party module if you’d like.
type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC822ZTime{t}
	return nil
}

func (rt RFC822ZTime) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	out := time.Now().Format(time.RFC3339)
	// out := rt.Time.Format(time.RFC822Z)
	w.Write([]byte(out))
}

// Tested using
// curl http://localhost:8080
func main() {
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      RFC822ZTime{},
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

// ex2:
// Write a small middleware component that uses JSON structured logging to log the IP address of each incoming request to your web server.

// ex3:
// Add the ability to return the time as JSON. Use the Accept header to control whether JSON or text is returned (default to text). The JSON should be structured as follows:
