package trace

type ExporterType int

const (
	Otlp ExporterType = iota
	Stdout
)

type Config struct {
	ReceiverUrl  string
	ServiceName  string
	Environment  string
	ExporterType ExporterType
}
