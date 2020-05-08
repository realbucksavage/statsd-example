package stats

import (
	"fmt"

	"gopkg.in/alexcesaro/statsd.v2"
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
	t.t.Send(t.ns + "." + n)
}

type statsdMetric struct {
	ns     string
	daemon *statsd.Client
}

func (s statsdMetric) Increment(n string) {
	s.daemon.Increment(s.ns + "." + n)
}

func (s statsdMetric) Gauge(n string, value interface{}) {
	s.daemon.Gauge(s.ns+"."+n, value)
}

func (s statsdMetric) TimingStart() Timer {
	return statsdTime{
		ns: s.ns,
		t:  s.daemon.NewTiming(),
	}
}

func NewCounter(ns string) (Metrics, error) {
	client, err := statsd.New()
	if err != nil {
		return nil, fmt.Errorf("statsd connect: %s", err)
	}
	sm := statsdMetric{
		ns:     ns,
		daemon: client,
	}

	return &sm, nil
}
