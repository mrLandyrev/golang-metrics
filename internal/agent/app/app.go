package app

import (
	"time"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/collect/exporters"
	collectService "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/collect/service"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/repository"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/sync/client"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/sync/service"
)

type Exporter interface {
	GetMetrics() ([]models.Metric, error)
}

type CollectService interface {
	RegisterExporter(exporter collectService.Exporter) error
	Collect() error
}

type SyncService interface {
	SyncMetrics() error
}

type App struct {
	collectService  CollectService
	syncService     SyncService
	syncInterval    int
	collectInterval int
}

func (app *App) Run() {
	var i int
	for i = 1; ; i++ {
		if (i % app.collectInterval) == 0 {
			app.collectService.Collect()
		}
		if (i % app.syncInterval) == 0 {
			app.syncService.SyncMetrics()
		}
		time.Sleep(time.Second)
	}
}

func NewApp(serverAddress string, syncInterval int, collectInterval int) *App {
	metricsRepository := repository.NewMemoryMetricsRepository()

	collectService := collectService.NewCollectService(metricsRepository)
	collectService.RegisterExporter(exporters.NewIncrementExproter())
	collectService.RegisterExporter(exporters.NewRandomExproter())
	collectService.RegisterExporter(exporters.NewRuntimeExporter())

	syncClient := client.NewHTTPClient(serverAddress)
	syncService := service.NewSyncService(metricsRepository, syncClient)

	return &App{syncService: syncService, collectService: collectService, syncInterval: syncInterval, collectInterval: collectInterval}
}
