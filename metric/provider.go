package metric

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/export/metric"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func createExporter(ctx context.Context, exporterType ExporterType, receiverUrl string) (sdkmetric.Exporter, error) {
	var exporter sdkmetric.Exporter
	var err error
	switch exporterType {
	case Otlp:
		exporter, err = otlpmetricgrpc.New(ctx,
			otlpmetricgrpc.WithInsecure(),
			otlpmetricgrpc.WithEndpoint("localhost:30080"),
		)
		if err != nil {
			return nil, err
		}
	case Stdout:
		exporter, err = stdoutmetric.New(
			stdoutmetric.WithPrettyPrint(),
		)
		if err != nil {
			return nil, err
		}
	}
	return exporter, nil
}

func InitProvider(config Config) (func() error, error) {
	ctx := context.Background()

	exporter, err := createExporter(ctx, config.ExporterType, config.ReceiverUrl)
	if err != nil {
		return nil, err
	}

	pusher := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			exporter,
		),
		controller.WithExporter(exporter),
		controller.WithCollectPeriod(config.CollectPeriod),
		controller.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
			attribute.String("environment", config.Environment),
		)),
	)
	err = pusher.Start(ctx)
	if err != nil {
		return nil, err
	}

	global.SetMeterProvider(pusher.MeterProvider())
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return func() error {
		return pusher.Stop(ctx)
	}, nil
}
