package tracing

type AnyWithTrace struct {
	Any     any    `json:"any,omitempty"`
	TraceID string `json:"uber_trace_id,omitempty" xml:"uber-trace-id"`
}
