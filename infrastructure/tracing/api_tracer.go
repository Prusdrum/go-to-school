package tracing

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Printer interface {
	
// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
Printf(format string, v ...any)
}

type Tracer struct {
	printer Printer
	handler http.Handler
}

func NewTracer(printer Printer, handler http.Handler) *Tracer {
	return &Tracer{printer, handler}
}

func (t *Tracer) trace(message string) {
	t.printer.Printf("TRACE: %s\n", message)
}

func (t*Tracer) ServeHTTP(w http.ResponseWriter ,req *http.Request) {
	requestUUID := uuid.New()
	t.trace(fmt.Sprintf( "[%s] %s %s - started", requestUUID.String(), req.Method, req.URL.Path))
	(t.handler).ServeHTTP(w, req)
	t.trace(fmt.Sprintf( "[%s] %s %s - finished", requestUUID.String(), req.Method, req.URL.Path))
}
