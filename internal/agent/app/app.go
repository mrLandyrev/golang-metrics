package app

import (
	"context"
	"fmt"
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
	cancel          context.CancelFunc
}

func (app *AgentApp) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			fmt.Println("sync")
			select {
			case <-ctx.Done():
				return
			default:
				_ = app.syncService.SyncMetrics()
			}

			time.Sleep(time.Duration(app.collectInterval) * time.Second)
		}
	}(ctx)
	go func(ctx context.Context) {
		for {
			fmt.Println("collect")
			select {
			case <-ctx.Done():
				return
			default:
				_ = app.collectService.Collect()
			}

			time.Sleep(time.Duration(app.syncInterval) * time.Second)
		}
	}(ctx)
	app.cancel = cancel
}

func (app *AgentApp) Stop() {
	app.cancel()
}

func NewAgentApp(config Config) *AgentApp {
	//build dependencies
	metricsRepository := storage.NewMemoryMetricsRepository()

	collectService := collectService.NewCollectService(metricsRepository)
	collectService.RegisterExporter(exporters.NewIncrementExproter())
	collectService.RegisterExporter(exporters.NewRandomExproter())
	collectService.RegisterExporter(exporters.NewRuntimeExporter())

	syncClient := rest.NewHTTPClient(config.ServerAddress, config.SignKey)
	syncService := service.NewSyncService(metricsRepository, syncClient)

	return &AgentApp{syncService: syncService, collectService: collectService, syncInterval: config.SyncInteval, collectInterval: config.CollectInterval}
}
