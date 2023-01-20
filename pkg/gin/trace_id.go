package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"log"
)

func LogTraceID(gtx *gin.Context) {
	gtx.Next()
	a := gtx.Request.Context().Value(tracing.HttpSpanKey)
	if a != nil {
		span, ok := a.(trace.Span)
		if ok {
			log.Printf("path: %s | traceID: %s", gtx.Request.URL, span.SpanContext().TraceID().String())
		}
	}
}

func TraceRequest(gtx *gin.Context) {
	ctx, span := tracing.Tracer.Start(gtx.Request.Context(), fmt.Sprintf("%s %s", gtx.Request.Method, gtx.FullPath()))
	span.SetAttributes(semconv.HTTPURLKey.String(gtx.Request.URL.String()))
	gtx.Request = gtx.Request.WithContext(context.WithValue(ctx, tracing.HttpSpanKey, span))

	gtx.Next()

	LogTraceID(gtx)
}
