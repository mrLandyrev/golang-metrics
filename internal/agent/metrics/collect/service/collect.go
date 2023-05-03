package service

import (
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type Exporter interface {
	GetMetrics() ([]models.Metric, error)
}

type MetricsRepository interface {
	GetByKindAndName(kind string, name string) (models.Metric, error)
	Persist(item models.Metric) error
}

type CollectService struct {
	exporters         []Exporter
	metricsRepository MetricsRepository
}

func (collectService *CollectService) RegisterExporter(exporter Exporter) error {
	collectService.exporters = append(collectService.exporters, exporter)
	return nil
}

func (collectService *CollectService) Collect() error {
	for _, exporter := range collectService.exporters {
		metrics, err := exporter.GetMetrics()

		if err != nil {
			return err
		}

		for _, metric := range metrics {
			record, err := collectService.metricsRepository.GetByKindAndName(metric.GetKind(), metric.GetName())

			if err != nil {
				return err
			}

			if record == nil {
				err = collectService.metricsRepository.Persist(metric)
				if err != nil {
					return err
				}
				continue
			}

			err = record.AddValue(metric.GetValue())

			if err != nil {
				return err
			}

			err = collectService.metricsRepository.Persist(record)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func NewCollectService(metricsRepository MetricsRepository) *CollectService {
	return &CollectService{exporters: []Exporter{}, metricsRepository: metricsRepository}
}
