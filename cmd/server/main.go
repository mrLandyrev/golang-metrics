package main

import (
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

func main() {
	metricsService := metrics.NewService(metrics.NewMemoryRepository(), metrics.NewFactory())

	server := metrics.NewServer(metricsService)
	server.Listen()
}
