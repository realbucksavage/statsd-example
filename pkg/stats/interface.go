package stats

import (
	"fmt"
	"io"

	"github.com/cactus/go-statsd-client/statsd"
)

type Metrics interface {
	Increment(string)
	Gauge(string, int64)
	TimeMillisecond(string, float32)

	io.Closer
}

type Timer interface {
	Send(string)
}

func NewCounter(ns string) (Metrics, error) {
	cfg := &statsd.ClientConfig{
		Address: "127.0.0.1:8125",
		Prefix:  ns,
	}
	client, err := statsd.NewClientWithConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("statsd connect: %s", err)
	}
	sm := statsdMetric{
		ns:     ns,
		daemon: client,
	}

	return &sm, nil
}
