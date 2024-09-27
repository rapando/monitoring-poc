package api

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of http requests",
		},
		[]string{"path"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of Http requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var response = map[string]any{
		"message": "ok",
		"time":    time.Now().Format(time.Kitchen),
		"success": true,
	}

	Response(w, http.StatusOK, response)
}
