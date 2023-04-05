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
	metricsFactory := factory.NewDefaultMetricsFactory()
	metricsService := service.NewDefaultMetricsService(metricsRepository, metricsFactory)

	router := router.NewRouter()
	router.Use("POST", "/update/:kind/:name/:value", handlers.GetUpdateMetricHandler(metricsService))
	router.Use("GET", "/get/:kind/:name", handlers.GetGetMetricHandler(metricsService))

	http.ListenAndServe(":8080", router)
}
