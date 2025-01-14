package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

// TracingMiddleware creates a middleware that starts a new trace span for each request.
// It uses the provided tracer name to create a new tracer instance.
// The middleware sets the context with the span to the next handler.
//
// Parameters:
//   - name: The name of the tracer.
//
// Returns:
//   - fiber.Handler: The middleware handler function.
func TracingMiddleware(name string) fiber.Handler {
	tracer := otel.Tracer(name)

	return func(c *fiber.Ctx) error {
		ctx, span := tracer.Start(c.Context(), c.Path())
		defer span.End()

		// Pass the context with the span to the next handler
		c.SetUserContext(ctx)
		return c.Next()
	}
}
