package main

import (
	"log"
	"net/http"
	"simpleWebService/pages"
	"simpleWebService/server"
)

var (
	ServerAddress      = ":8443"
	CertificateFile    = "certs/localhost.crt"
	CertificateKeyFile = "certs/localhost.key"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pages.HomeHandler)

	srv := server.New(mux, ServerAddress)

	err := srv.ListenAndServeTLS(CertificateFile, CertificateKeyFile)
	if err != nil {
		log.Fatalf("Error during server startup: %v", err)
	}
}
