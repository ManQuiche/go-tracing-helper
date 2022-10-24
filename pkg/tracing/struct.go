package tracing

type AnyWithTrace[T any] struct {
	Any     T      `json:"any,omitempty"`
	TraceID string `json:"uber_trace_id,omitempty" xml:"uber-trace-id"`
}
