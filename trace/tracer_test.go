package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return of 'New' is nil")
	} else {
		trace.Trace("Hello, trace package")
		if buf.String() != "Hello, trace package\n" {
			t.Errorf("Wrong chars '%s' were output.", buf.String())
		}
	}
}
