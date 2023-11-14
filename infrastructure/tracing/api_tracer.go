package tracing

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
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

func (t *Tracer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	requestUUID := uuid.New()
	t.trace(fmt.Sprintf("[%s] %s %s - started", requestUUID.String(), req.Method, req.URL.Path))
	context := req.Context()

	traceLogger := &TraceLogger{t.printer, requestUUID}

	newContext := newContextWithTracingLogger(context, traceLogger)

	req = req.WithContext(newContext)

	(t.handler).ServeHTTP(w, req)
	t.trace(fmt.Sprintf("[%s] %s %s - finished", requestUUID.String(), req.Method, req.URL.Path))
}

type TraceLogger struct {
	printer Printer
	id      uuid.UUID
}

func newContextWithTracingLogger(ctx context.Context, logger *TraceLogger) context.Context {
	return context.WithValue(ctx, "traceLogger", logger)
}

func TraceLoggerFromContext(ctx context.Context) (*TraceLogger, bool) {
	logger, ok := ctx.Value("traceLogger").(*TraceLogger)
	return logger, ok
}

func (t *TraceLogger) Printf(format string, v ...any) {
	t.printer.Printf("TRACE: [%s] %s", t.id.String(), fmt.Sprintf(format, v...))
}

func (t *TraceLogger) Println(v ...any) {
	t.Printf("%s", v)
}
