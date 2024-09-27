package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rapando/monitoring-poc/src/pkg/log"
)

type ContextKey string

var (
	RequestIDKey ContextKey = "request_id"
)

func RequestIDMW(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var start = time.Now()
		var _uuid = uuid.New().String()
		log.Infof(_uuid, "[%s] %s", r.Method, r.URL.Path)
		var ctx = context.WithValue(r.Context(), RequestIDKey, _uuid)
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Infof(_uuid, "[%s] %s took %s", r.Method, r.URL.Path, time.Since(start))
		requestCount.WithLabelValues(r.URL.Path).Inc()
		var duration = time.Now().Sub(start).Microseconds()
		requestDuration.WithLabelValues(r.URL.Path).Observe(float64(duration))
	})
}

func Response(w http.ResponseWriter, code int, response any) {
	resBytes, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("connection", "close")
	w.WriteHeader(code)
	_, _ = w.Write(resBytes)
}
