package pages

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello, world!"

type Handler struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Request processed")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message + " " + time.Now().Format("2006-01-02 15:04:05.999999999")))
	if err != nil {
		h.logger.Fatalf("Error writing the response: %v", err)
	}
}
