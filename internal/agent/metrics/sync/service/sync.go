package service

import "github.com/mrLandyrev/golang-metrics/internal/metrics"

type Client interface {
	SendMetric(metric metrics.Metric) error
	SendMetrics(metrics []metrics.Metric) error
}

type MetricsRepository interface {
	GetAll() ([]metrics.Metric, error)
	Clear() error
}

type SyncService struct {
	metricsRepository MetricsRepository
	syncClient        Client
}

func (syncService *SyncService) SyncMetrics() error {
	metrics, err := syncService.metricsRepository.GetAll()

	if err != nil {
		return err
	}

	err = syncService.syncClient.SendMetrics(metrics)

	if err != nil {
		return err
	}

	return syncService.metricsRepository.Clear()
}

func NewSyncService(metricsRepository MetricsRepository, syncClient Client) *SyncService {
	return &SyncService{metricsRepository: metricsRepository, syncClient: syncClient}
}
