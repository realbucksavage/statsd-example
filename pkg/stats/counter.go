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
	s.daemon.Incr(n, 1)
}

func (s statsdMetric) Gauge(n string, value int64) {
	log.Printf("Gauge %s.%s at %v", s.ns, n, value)
	s.daemon.Gauge(n, value)
}

func (s statsdMetric) Time(n string, t time.Duration) {
	log.Printf("Logged %s.%s at %dms", s.ns, n, t.Microseconds())
	s.daemon.PrecisionTiming(n, t)
}

func (s statsdMetric) Close() error {
	return s.daemon.Close()
}
