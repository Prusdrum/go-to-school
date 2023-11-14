package tracing

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// / Provides a way to print messages.
type PrintProvider interface {

	// Printf formats according to a format specifier and writes to standard output.
	// Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...any)
}

type TracingHandler struct {
	printer PrintProvider
	handler http.Handler
}

func NewTracingHandler(printer PrintProvider, handler http.Handler) http.Handler {
	return &TracingHandler{printer, handler}
}

func (t *TracingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	requestUUID := uuid.New()

	traceLogger := &TraceLogger{t.printer, requestUUID}
	traceLogger.Printf("%s %s - started", req.Method, req.URL.Path)
	context := req.Context()

	newContext := newContextWithTracingLogger(context, traceLogger)

	req = req.WithContext(newContext)

	(t.handler).ServeHTTP(w, req)

	traceLogger.Printf("%s %s - finished", req.Method, req.URL.Path)
}

func newContextWithTracingLogger(ctx context.Context, logger *TraceLogger) context.Context {
	return context.WithValue(ctx, "traceLogger", logger)
}

func TraceLoggerFromRequest(req *http.Request) (*TraceLogger, bool) {
	logger, ok := req.Context().Value("traceLogger").(*TraceLogger)
	return logger, ok
}
