package pages

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello, world!"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request processed")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message + " " + time.Now().Format("2006-01-02 15:04:05.999999999")))
	if err != nil {
		log.Fatalf("Error writing the response: %v", err)
	}
}
