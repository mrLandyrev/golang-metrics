package service

import "github.com/mrLandyrev/golang-metrics/internal/metrics"

type Client interface {
	SendMetric(metric metrics.Metric) error
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

	for _, metric := range metrics {
		err = syncService.syncClient.SendMetric(metric)

		if err != nil {
			return err
		}
	}

	return syncService.metricsRepository.Clear()
}

func NewSyncService(metricsRepository MetricsRepository, syncClient Client) *SyncService {
	return &SyncService{metricsRepository: metricsRepository, syncClient: syncClient}
}
