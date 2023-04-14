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
	Collect() ([]models.Metric, error)
}

type CollectService interface {
	RegisterExporter(exporter collectService.Exporter) error
	Collect() error
}

type SyncService interface {
	SyncMetrics() error
}

type App struct {
	collectService CollectService
	syncService    SyncService
}

func (app *App) Run() {
	for i := 1; ; i++ {
		app.collectService.Collect()
		if i%5 == 0 {
			app.syncService.SyncMetrics()
		}
		time.Sleep(time.Second * 2)
	}
}

func NewApp() *App {
	metricsRepository := repository.NewMemoryMetricsRepository()

	collectService := collectService.NewCollectService(metricsRepository)
	collectService.RegisterExporter(exporters.NewIncrementExproter())
	collectService.RegisterExporter(exporters.NewRandomExproter())
	collectService.RegisterExporter(exporters.NewRuntimeExporter())

	syncClient := client.NewHttpClient()
	syncService := service.NewSyncService(metricsRepository, syncClient)

	return &App{syncService: syncService, collectService: collectService}
}
