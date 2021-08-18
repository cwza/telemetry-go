package trace

import (
	"context"
	"fmt"
	"testing"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("tracer")

func foo1(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "foo1")
	defer span.End()
	foo2(ctx)
}

func foo2(ctx context.Context) {
	_, span := tracer.Start(ctx, "foo2")
	defer span.End()
}

func bar1(ctx context.Context) {
	_, span := tracer.Start(ctx, "bar1")
	span.RecordError(fmt.Errorf("bar1 error"))
	defer span.End()
}

func run() {
	ctx, span := tracer.Start(context.Background(), "run")
	// This is log, which will automatically add timestamp for you
	span.AddEvent("event1", trace.WithAttributes(attribute.Int("int1", 100)))
	// This is tag, can be query by ui
	span.SetAttributes(attribute.String("key1", "value1"))
	defer span.End()
	foo1(ctx)
	bar1(ctx)
}

func TestInitProvider_Stdout(t *testing.T) {
	config := Config{
		ReceiverUrl:  "localhost:30080",
		ServiceName:  "test-service",
		Environment:  "test",
		ExporterType: Stdout,
	}
	shutdown, err := InitProvider(config)
	if err != nil {
		t.Errorf("failed to InitProvider, %s", err)
	}
	defer func() {
		err := shutdown()
		if err != nil {
			t.Errorf("failed to shutdown, %s", err)
		}
	}()

	run()
}

func TestInitProvider_Otlp(t *testing.T) {
	config := Config{
		ReceiverUrl:  "localhost:30080",
		ServiceName:  "test-service",
		Environment:  "test",
		ExporterType: Otlp,
	}
	shutdown, err := InitProvider(config)
	if err != nil {
		t.Errorf("failed to InitProvider, %s", err)
	}
	defer func() {
		err := shutdown()
		if err != nil {
			t.Errorf("failed to shutdown, %s", err)
		}
	}()

	run()
}
