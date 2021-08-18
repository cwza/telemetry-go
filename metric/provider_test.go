package metric

import (
	"context"
	"log"
	"testing"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

const (
	ReceiverUrl = "localhost:30080"
	ServiceName = "test-service"
	Environment = "test"
)

var (
	meter          = global.Meter("meter")
	boundedCounter metric.BoundInt64Counter
	boundedValue   metric.BoundInt64ValueRecorder
)

func init() {
	counter, err := meter.NewInt64Counter("test-counter", metric.WithDescription("counter test"))
	if err != nil {
		log.Fatalf("failed to NewInt64Counter")
	}
	boundedCounter = counter.Bind(attribute.String("key1", "value1")) // bind labels to counter
	value, err := meter.NewInt64ValueRecorder("test-value", metric.WithDescription("value test"))
	if err != nil {
		log.Fatalf("failed to NewInt64ValueRecorder")
	}
	boundedValue = value.Bind(attribute.String("key1", "value1"))
}

func run() {
	boundedCounter.Add(context.Background(), 1)
	boundedCounter.Add(context.Background(), 1)
	boundedCounter.Add(context.Background(), 1)
	boundedValue.Record(context.Background(), 23)
	boundedValue.Record(context.Background(), 25)
	boundedValue.Record(context.Background(), 30)

	boundedCounter.Unbind()
	boundedValue.Unbind()
}

func TestInitProvider_Stdout(t *testing.T) {
	config := Config{
		ReceiverUrl:   "",
		ServiceName:   "test-service",
		Environment:   "test",
		CollectPeriod: 5 * time.Second,
		ExporterType:  Stdout,
	}
	shutdown, err := InitProvider(config)
	if err != nil {
		log.Fatalf("failed to InitProvider, %s", err)
	}
	defer func() {
		err := shutdown()
		if err != nil {
			log.Fatalf("failed to shutdown, %s", err)
		}
	}()

	run()
}

func TestInitProvider_Otlp(t *testing.T) {
	config := Config{
		ReceiverUrl:   "localhost:30080",
		ServiceName:   "test-service",
		Environment:   "test",
		CollectPeriod: 5 * time.Second,
		ExporterType:  Otlp,
	}
	shutdown, err := InitProvider(config)
	if err != nil {
		log.Fatalf("failed to InitProvider, %s", err)
	}
	defer func() {
		err := shutdown()
		if err != nil {
			log.Fatalf("failed to shutdown, %s", err)
		}
	}()

	run()
}
