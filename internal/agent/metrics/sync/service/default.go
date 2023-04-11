package service

import (
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/repository"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/sync/client"
)

type DefaultSyncService struct {
	metricsRepository repository.MetricsRepository
	syncClient        client.Client
}

func (syncService *DefaultSyncService) SyncMetrics() error {
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

	return nil
}

func NewDefaultSyncService(metricsRepository repository.MetricsRepository, syncClient client.Client) *DefaultSyncService {
	return &DefaultSyncService{metricsRepository: metricsRepository, syncClient: syncClient}
}
