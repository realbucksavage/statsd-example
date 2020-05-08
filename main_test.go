package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/realbucksavage/statsd-example/pkg/api"
	"github.com/realbucksavage/statsd-example/pkg/stats"
)

func TestGauge(t *testing.T) {
	m, err := stats.NewCounter("api_test.suite")
	if err != nil {
		t.Fatal(err)
	}

	router := api.NewRouter(m)
	server := httptest.NewServer(router)
	defer server.Close()

	client := &http.Client{}

	t.Log("Creating 20 GoRoutines data per second.")
	go func() {
		for {
			if _, err := client.Get("http://" + server.Listener.Addr().String()); err != nil {
				t.Fatal(err)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	t.Log("Test complete.")
}
