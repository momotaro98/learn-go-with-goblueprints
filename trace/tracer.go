package trace

import (
	"fmt"
	"io"
)

// Tracer is an interface コード内でのできごとを記録できるオブジェクトを表す
type Tracer interface {
	Trace(...interface{})
}

// tracer is a struct that implements Tracer
type tracer struct {
	out io.Writer
}

// New is a constructor of Tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Trace Writes to out-stream
func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// nilTracer implements Tracer that does nothing
type nilTracer struct{}

// Off returns nilTracer
func Off() Tracer {
	return &nilTracer{}
}

// Trace is a method nilTracer has
func (t *nilTracer) Trace(a ...interface{}) {}
