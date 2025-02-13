package tracing

import (
	"time"

	"github.com/cnosdb/db/pkg/tracing/fields"
	"github.com/cnosdb/db/pkg/tracing/labels"
)

// RawSpan represents the data associated with a span.
type RawSpan struct {
	Context      SpanContext
	ParentSpanID uint64        // ParentSpanID identifies the parent of this span or 0 if this is the root span.
	Name         string        // Name is the operation name given to this span.
	Start        time.Time     // Start identifies the start time of the span.
	Labels       labels.Labels // Labels contains additional metadata about this span.
	Fields       fields.Fields // Fields contains typed values associated with this span.
}
