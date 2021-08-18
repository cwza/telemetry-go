module github.com/cwza/telemetry-go

go 1.16

require (
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.22.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.22.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0-RC2
	go.opentelemetry.io/otel/metric v0.22.0
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
	go.opentelemetry.io/otel/sdk/export/metric v0.22.0
	go.opentelemetry.io/otel/sdk/metric v0.22.0
	go.opentelemetry.io/otel/trace v1.0.0-RC2
)
