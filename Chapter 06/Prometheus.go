package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create Prometheus metrics
	requestCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "status"},
	)
	responseTime := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_time_seconds",
			Help:    "HTTP response time distribution",
			Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
		},
		[]string{"method", "status"},
	)

	// Register metrics with Prometheus
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(responseTime)

	// Create an HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment the request count
		requestCount.WithLabelValues(r.Method, "200").Inc()

		// Simulate processing time
		// In a real application, this would be replaced with actual request processing logic
		time.Sleep(100 * time.Millisecond)

		// Record the response time
		responseTime.WithLabelValues(r.Method, "200").Observe(time.Since(start).Seconds())

		// Send response
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello, World!")
	})

	// Expose metrics and application endpoints
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}