package stats

import (
	"log"
	"time"

	"github.com/cactus/go-statsd-client/statsd"
)

type statsdMetric struct {
	ns     string
	daemon statsd.Statter
}

func (s statsdMetric) Increment(n string) {
	log.Printf("Increment %s.%s", s.ns, n)
	s.daemon.Inc(n, 1, 1.0)
}

func (s statsdMetric) Gauge(n string, value int64) {
	log.Printf("Gauge %s.%s at %v", s.ns, n, value)
	s.daemon.Gauge(n, value, 1.0)
}

func (s statsdMetric) TimeMillisecond(n string, t float32) {
	log.Printf("Logged %s.%s at %fms", s.ns, n, t)
	s.daemon.TimingDuration(n, time.Millisecond*time.Duration(t), 1.0)
}

func (s statsdMetric) Close() error {
	return s.daemon.Close()
}
