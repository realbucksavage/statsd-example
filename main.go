package main

import (
	"net/http"

	"github.com/prometheus/common/log"
	"github.com/realbucksavage/statsd-example/pkg/api"
	"github.com/realbucksavage/statsd-example/pkg/stats"
)

func main() {
	m, err := stats.NewCounter("testapp")
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", api.NewRouter(m))
}
