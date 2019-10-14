package main

import (
	"log"
	"net/http"
)

const message = "Hello, world!"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(message))
		if err != nil {
			log.Fatalf("Error writing the response: %v", err)
		}
	})

	err := http.ListenAndServe(":8088", mux)
	if err != nil {
		log.Fatalf("Error during server startup: %v", err)
	}
}
