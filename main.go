package main

import (
	"log"
	"runtime"
	"time"

	"gopkg.in/alexcesaro/statsd.v2"
)

func main() {
	c, err := statsd.New()
	if err != nil {
		log.Fatalf("start statsd client: %s", err)
	}
	defer c.Close()

	c.Increment("testapp.counter")
	c.Gauge("testapp.goroutines", runtime.NumGoroutine())

	t := c.NewTiming()
	time.Sleep(500 * time.Millisecond)
	t.Send("testapp.latency")
}
