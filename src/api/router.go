package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRouter(router *chi.Mux) {
	router.Get("/", RequestIDMW(HomeHandler))
	router.Handle("/metrics", promhttp.Handler())
	router.Post("/data", RequestIDMW(AddHandler))
	router.Get("/data", RequestIDMW(CountHandler))
}
