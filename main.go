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

	p := pages.New(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/", p.Home)

	srv := server.New(mux, ServerAddress)

	err := srv.ListenAndServeTLS(CertificateFile, CertificateKeyFile)
	if err != nil {
		log.Fatalf("Error during server startup: %v", err)
	}
}
