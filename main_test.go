package main

import (
	"net/http"
	"net/http/httptest"
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
	for i := 0; i < 20; i++ {
		if _, err := client.Get("http://" + server.Listener.Addr().String()); err != nil {
			t.Fatal(err)
		}

		time.Sleep(1 * time.Second)
	}
}
