package stats

import (
	"fmt"
	"io"
	"time"

	"github.com/quipo/statsd"
)

type Metrics interface {
	Increment(string)
	Gauge(string, int64)
	Time(string, time.Duration)

	io.Closer
}

type Timer interface {
	Send(string)
}

func NewCounter(ns string) (Metrics, error) {
	client := statsd.NewStatsdClient("127.0.0.1:8125", ns)
	if err := client.CreateSocket(); err != nil {
		return nil, fmt.Errorf("statsd connect: %s", err)
	}

	sm := statsdMetric{
		ns:     ns,
		daemon: client,
	}

	return &sm, nil
}
