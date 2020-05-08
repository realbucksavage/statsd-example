package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/realbucksavage/statsd-example/pkg/api"
	"github.com/realbucksavage/statsd-example/pkg/stats"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	m, err := stats.NewCounter("testapi")
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	quit := make(chan os.Signal)
	done := make(chan bool, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	log.Println("Starting API server at :8080")
	server := newApiServer(m)
	go shutdown(server, quit, done)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server startup failed: %s", err)
	}

	<-done
	log.Println("Shutdown complete")
}

func shutdown(server *http.Server, quit chan os.Signal, done chan bool) {
	<-quit
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %s", err)
	}

	close(done)
}

func newApiServer(m stats.Metrics) *http.Server {
	router := api.NewRouter(m)
	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
