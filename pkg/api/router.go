package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/realbucksavage/statsd-example/pkg/stats"
)

func NewRouter(m stats.Metrics) http.Handler {
	r := mux.NewRouter()
	r.Handle("/", goroutines(m))
	r.Handle("/random", randomNum(m))

	return r
}
