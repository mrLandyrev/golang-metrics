package service

import (
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

// -- dependencies --

type Exporter interface {
	GetMetrics() ([]metrics.Metric, error)
}

type MetricsRepository interface {
	GetByKindAndName(kind string, name string) (metrics.Metric, error)
	Persist(item metrics.Metric) error
}

// -- dependencies --

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
			record, err := collectService.metricsRepository.GetByKindAndName(metric.Kind(), metric.Name())

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

			err = record.AddValue(metric.Value())

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
