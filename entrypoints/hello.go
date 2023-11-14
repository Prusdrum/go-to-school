package entrypoints

import (
	"fmt"
	"go-to-school/main/infrastructure/tracing"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

// DI
func NewHello(log *log.Logger) *Hello {
	return &Hello{log}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log, _ := tracing.TraceLoggerFromContext(req.Context())
	log.Println("Request received")
	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "Error ocurred", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Hello %s", body)
}
