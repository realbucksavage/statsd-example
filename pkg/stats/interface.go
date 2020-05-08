package stats

import (
	"fmt"

	"gopkg.in/alexcesaro/statsd.v2"
	"log"
)

type Metrics interface {
	Increment(string)
	Gauge(string, interface{})
	TimingStart() Timer
}

type Timer interface {
	Send(string)
}

type statsdTime struct {
	ns string
	t  statsd.Timing
}

func (t statsdTime) Send(n string) {
	fmt.Printf("Time %s.%s at %dms", t.ns, n, t.t.Duration().Microseconds())
	t.t.Send(n)
}

type statsdMetric struct {
	ns     string
	daemon *statsd.Client
}

func (s statsdMetric) Increment(n string) {
	log.Printf("Increment %s.%s", s.ns, n)
	s.daemon.Increment(n)
}

func (s statsdMetric) Gauge(n string, value interface{}) {
	log.Printf("Gauge %s.%s at %v", s.ns, n, value)
	s.daemon.Gauge(n, value)
}

func (s statsdMetric) TimingStart() Timer {
	return statsdTime{
		ns: s.ns,
		t:  s.daemon.NewTiming(),
	}
}

func NewCounter(ns string) (Metrics, error) {
	client, err := statsd.New(
		statsd.Prefix(ns),
	)
	if err != nil {
		return nil, fmt.Errorf("statsd connect: %s", err)
	}
	sm := statsdMetric{
		ns:     ns,
		daemon: client,
	}

	return &sm, nil
}
