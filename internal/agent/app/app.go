package app

import (
	"time"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/collect/exporters"
	collectService "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/collect/service"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/sync/service"
	"github.com/mrLandyrev/golang-metrics/internal/metrics/storage"
	"github.com/mrLandyrev/golang-metrics/internal/server/app/transport/rest"
)

type CollectService interface {
	Collect() error
}

type SyncService interface {
	SyncMetrics() error
}

type AgentApp struct {
	collectService  CollectService
	syncService     SyncService
	syncInterval    int
	collectInterval int
}

func (app *AgentApp) Run() {
	for i := 1; ; i++ {
		if (i % app.collectInterval) == 0 {
			_ = app.collectService.Collect()
		}
		if (i % app.syncInterval) == 0 {
			_ = app.syncService.SyncMetrics()
		}
		time.Sleep(time.Second)
	}
}

func NewAgentApp(config Config) *AgentApp {
	//build dependencies
	metricsRepository := storage.NewMemoryMetricsRepository()

	collectService := collectService.NewCollectService(metricsRepository)
	collectService.RegisterExporter(exporters.NewIncrementExproter())
	collectService.RegisterExporter(exporters.NewRandomExproter())
	collectService.RegisterExporter(exporters.NewRuntimeExporter())

	syncClient := rest.NewHTTPClient(config.ServerAddress)
	syncService := service.NewSyncService(metricsRepository, syncClient)

	return &AgentApp{syncService: syncService, collectService: collectService, syncInterval: config.SyncInteval, collectInterval: config.CollectInterval}
}
