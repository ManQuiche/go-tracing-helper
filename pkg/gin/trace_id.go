package gin

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	"go.opentelemetry.io/otel/trace"
	"log"
)

func LogTraceID(gtx *gin.Context) {
	gtx.Next()
	a := gtx.Request.Context().Value(tracing.HttpSpanKey)

	span, ok := a.(trace.Span)
	if ok {
		log.Printf("path: %s | traceID: %s", gtx.Request.URL, span.SpanContext().TraceID().String())
	}
}
