package main

import (
	"log"
	"net/http"
	"simpleWebService/server"
	"time"
)

const message = "Hello, world!"

var (
	ServerAddress      = ":8443"
	CertificateFile    = "certs/localhost.crt"
	CertificateKeyFile = "certs/localhost.key"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(message + " " + time.Now().Format("2006-01-02 15:04:05.999999999")))
		if err != nil {
			log.Fatalf("Error writing the response: %v", err)
		}
	})

	srv := server.New(mux, ServerAddress)

	err := srv.ListenAndServeTLS(CertificateFile, CertificateKeyFile)
	if err != nil {
		log.Fatalf("Error during server startup: %v", err)
	}
}
