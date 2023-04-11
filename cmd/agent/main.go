package main

import (
	"time"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/collect/exporers"
	colectService "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/collect/service"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/repository"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/sync/client"
	syncService "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/sync/service"
)

func main() {
	metricsRepository := repository.NewMemoryMetricsRepository()

	collectService := colectService.NewDefaultCollectService(metricsRepository)
	collectService.RegisterExporter(exporers.NewRuntimeExporter())
	collectService.RegisterExporter(exporers.NewIncrementExproter())
	collectService.RegisterExporter(exporers.NewRandomExproter())

	syncClient := client.NewFmtClient()
	syncService := syncService.NewDefaultSyncService(metricsRepository, syncClient)

	for i := 0; ; i++ {
		collectService.Collect()

		if i%5 == 0 {
			err := syncService.SyncMetrics()
			if err == nil {
				metricsRepository.Clear()
			}
		}
		time.Sleep(time.Second * 2)
	}
}
