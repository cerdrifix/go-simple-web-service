package main

import (
	"log"
	"net/http"
	"os"
	"simpleWebService/pages"
	"simpleWebService/server"
)

var (
	ServerAddress      = ":8443"
	CertificateFile    = "certs/localhost.crt"
	CertificateKeyFile = "certs/localhost.key"
)

func main() {

	logger := log.New(os.Stdout, "cfx | ", log.Lshortfile|log.LstdFlags)

	mux := http.NewServeMux()
	setupRoutes(mux, pages.New(logger))

	srv := server.New(mux, ServerAddress)

	err := srv.ListenAndServeTLS(CertificateFile, CertificateKeyFile)
	if err != nil {
		log.Fatalf("Error during server startup: %v", err)
	}
}

func setupRoutes(mux *http.ServeMux, h *pages.Handler) {
	mux.HandleFunc("/", h.Logger(h.Home))
}
