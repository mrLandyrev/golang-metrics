package service

import (
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/collect/exporers"
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/repository"
)

type DefaultCollectService struct {
	exporters         []exporers.Exporter
	metricsRepository repository.MetricsRepository
}

func (collectService *DefaultCollectService) RegisterExporter(exporter exporers.Exporter) {
	collectService.exporters = append(collectService.exporters, exporter)
}

func (collectService *DefaultCollectService) Collect() error {
	for _, exporter := range collectService.exporters {
		err := exporter.Collect()
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDefaultCollectService(metricsRepository repository.MetricsRepository) *DefaultCollectService {
	return &DefaultCollectService{exporters: []exporers.Exporter{}, metricsRepository: metricsRepository}
}
