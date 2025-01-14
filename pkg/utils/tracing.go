package utils

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

// InitTracer initializes a new OpenTelemetry tracer with a stdout exporter
// that prints trace data in a human-readable format. It returns a function
// that can be called to shut down the tracer provider and clean up resources.
//
// Usage:
//
//	cleanup := InitTracer()
//	defer cleanup()
//
// Returns:
//
//	A function that shuts down the tracer provider when called.
func InitTracer() func() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize stdouttrace exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)

	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %v", err)
		}
	}
}
