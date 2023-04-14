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
	r              int64
	p              int64
}

func (app *App) Run() {
	var i int64
	for i = 1; ; i++ {
		if i%app.p == 0 {
			app.collectService.Collect()
		}
		if i%app.r == 0 {
			app.syncService.SyncMetrics()
		}
		time.Sleep(time.Second)
	}
}

func NewApp(a string, r int64, p int64) *App {
	metricsRepository := repository.NewMemoryMetricsRepository()

	collectService := collectService.NewCollectService(metricsRepository)
	collectService.RegisterExporter(exporters.NewIncrementExproter())
	collectService.RegisterExporter(exporters.NewRandomExproter())
	collectService.RegisterExporter(exporters.NewRuntimeExporter())

	syncClient := client.NewHTTPClient(a)
	syncService := service.NewSyncService(metricsRepository, syncClient)

	return &App{syncService: syncService, collectService: collectService}
}
