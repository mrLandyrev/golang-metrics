package main

import (
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/server/handlers"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/factory"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/repository"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
	"github.com/mrLandyrev/golang-metrics/pkg/router"
)

func main() {
	metricsRepository := repository.NewMemoryMetricsRepository()
	metricsFactory := factory.NewMetricsFactory()
	metricsService := service.NewMetricsService(metricsRepository, metricsFactory)

	router := router.NewRouter()
	router.Use("POST", "/update/:kind/:name/:value", handlers.BuildUpdateMetricHandler(metricsService))
	router.Use("GET", "/get/:kind/:name", handlers.BuildGetMetricHandler(metricsService))

	http.ListenAndServe(":8080", router)
}
