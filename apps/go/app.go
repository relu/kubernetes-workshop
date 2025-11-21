package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const PORT = 3000

func main() {
	_, name, _, _ := runtime.Caller(0)
	name = filepath.Base(name)
	if value, ok := os.LookupEnv("NAME"); ok {
		name = value
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s [%s] \"%s %s HTTP/1.1\" %d %s\n",
			r.RemoteAddr,
			time.Now().Format(time.RFC3339),
			r.Method,
			r.URL.Path,
			http.StatusOK,
			r.UserAgent())
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from %s", name)
	})

	log.Printf("listening on %d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), mux); err != nil {
		log.Fatal(err)
	}
}
