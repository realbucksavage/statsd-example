package stats

import (
	"log"
	"time"

	"github.com/quipo/statsd"
)

type statsdMetric struct {
	ns     string
	daemon *statsd.StatsdClient
}

func (s statsdMetric) Increment(n string) {
	log.Printf("Increment %s.%s", s.ns, n)
	if err := s.daemon.Incr(n, 1); err != nil {
		log.Fatalf("increment op: %s", err)
	}
}

func (s statsdMetric) Gauge(n string, value int64) {
	log.Printf("Gauge %s.%s at %v", s.ns, n, value)
	if err := s.daemon.Gauge(n, value); err != nil {
		log.Fatalf("gauge op: %s", err)
	}
}

func (s statsdMetric) Time(n string, t time.Duration) {
	log.Printf("Logged %s.%s at %dms", s.ns, n, t.Milliseconds())
	if err := s.daemon.PrecisionTiming(n, t); err != nil {
		log.Fatalf("time op: %s", err)
	}
}

func (s statsdMetric) Close() error {
	return s.daemon.Close()
}
