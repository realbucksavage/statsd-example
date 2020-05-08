package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/realbucksavage/statsd-example/pkg/stats"
)

func randomNum(m stats.Metrics) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		t := m.TimingStart()
		defer t.Send("latency")

		// Create fake latency
		n := rand.Intn(500)
		time.Sleep(time.Millisecond * time.Duration(n))

		resp := fmt.Sprintf(`{"num": %d}`, rand.Intn(100))

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(resp))

		m.Increment("request_count_random")
	}
}

func goroutines(m stats.Metrics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gr := runtime.NumGoroutine()
		resp := fmt.Sprintf(`{"num": %d}`, gr)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(resp))

		m.Gauge("request_count_gr", gr)
	}
}
