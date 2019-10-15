package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	db, err := sqlx.Open("postgres", "postgres://cerdrifix:cerdrifix1234@localhost:5432/tests?sslmode=disable")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}

	// Executing a ping to check if we can connect
	err = db.Ping()
	if err != nil {
		logger.Fatalf("Error. Can't connect to database: %v", err)
	}

	h := pages.New(logger, db)

	mux := http.NewServeMux()
	setupRoutes(mux, h)

	srv := server.New(mux, ServerAddress)

	err = srv.ListenAndServeTLS(CertificateFile, CertificateKeyFile)
	if err != nil {
		log.Fatalf("Error during server startup: %v", err)
	}
}

func setupRoutes(mux *http.ServeMux, h *pages.Handler) {
	mux.HandleFunc("/", h.Logger(h.Home))
}
