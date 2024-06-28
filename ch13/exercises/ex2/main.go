package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// ex2:
// Write a small middleware component that uses JSON structured logging to log the IP address of each incoming request to your web server.
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

func CreateSlogger() *slog.Logger {
	options := &slog.HandlerOptions{}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)
	return (mySlog)
}

// Tested using
// curl http://localhost:8080
func main() {

	mySlogger := CreateSlogger()
	serverLogger := ServerLogger(mySlogger)

	mux := http.NewServeMux()
	mux.Handle("/json", serverLogger(RFC822ZTime{}))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      serverLogger(RFC822ZTime{}),
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

func JsonLogging(r *http.Request) {
	data := jsonLog{
		IpAddr: r.RemoteAddr,
		Time:   time.Now(),
	}
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}

func UsingSlogger(logger *slog.Logger, r *http.Request) {
	logger.Info("ip address", "ip", r.RemoteAddr)
}

func ServerLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("Json logging")
			JsonLogging(r)
			fmt.Println("Slogger logging")
			UsingSlogger(logger, r)
			h.ServeHTTP(w, r)
		})
	}
}
