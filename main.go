package main

import (
	"net/http"

	"github.com/realbucksavage/statsd-example/pkg/api"
	"github.com/realbucksavage/statsd-example/pkg/stats"

	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	m, err := stats.NewCounter("testapi")
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", api.NewRouter(m))
}
