package main

import (
	"fmt"
	"net/http"
)

const message = "Hello, world!"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error writing the response: %v", err)
		}
	})

	err := http.ListenAndServe(":8088", mux)
	if err != nil {
		fmt.Printf("Error during server startup: %v", err)
	}
}
