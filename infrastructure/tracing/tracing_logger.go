package tracing

import (
	"fmt"
	"github.com/google/uuid"
)

type TraceLogger struct {
	printer PrintProvider
	id      uuid.UUID
}

func (t *TraceLogger) Printf(format string, v ...any) {
	t.printer.Printf("TRACE: [%s] %s", t.id.String(), fmt.Sprintf(format, v...))
}

func (t *TraceLogger) Println(v ...any) {
	t.Printf("%s\n", v...)
}
