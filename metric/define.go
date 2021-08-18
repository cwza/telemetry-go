package metric

import "time"

type ExporterType int

const (
	Otlp ExporterType = iota
	Stdout
)

type Config struct {
	ReceiverUrl   string
	ServiceName   string
	Environment   string
	CollectPeriod time.Duration
	ExporterType  ExporterType
}
