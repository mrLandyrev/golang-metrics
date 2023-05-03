package service

import (
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

// -- dependencies --

type MetricsFactory interface {
	GetInstance(kind string, name string) (metrics.Metric, error)
}

type MetricsRepository interface {
	GetAll() ([]metrics.Metric, error)
	GetByKindAndName(kind string, name string) (metrics.Metric, error)
	Persist(metric metrics.Metric) error
}

// -- dependencies --

type MetricsService struct {
	metricsRepository MetricsRepository
	metricsFactory    MetricsFactory
}

func (service *MetricsService) GetAll() ([]metrics.Metric, error) {
	return service.metricsRepository.GetAll()
}

func (service *MetricsService) AddRecord(kind string, name string, value string) error {

	// find item in storage
	item, err := service.metricsRepository.GetByKindAndName(kind, name)

	if err != nil {
		return err
	}

	// create item if not found in storage
	if item == nil {
		item, err = service.metricsFactory.GetInstance(kind, name)

		if err != nil {
			return err
		}
	}

	// apply raw value (how to interpret value encapsulated in metric)
	err = item.AddValue(value)

	if err != nil {
		return err
	}

	// save to storage
	return service.metricsRepository.Persist(item)
}

func (service *MetricsService) GetRecord(kind string, name string) (metrics.Metric, error) {
	return service.metricsRepository.GetByKindAndName(kind, name)
}

func NewMetricsService(metricsRepository MetricsRepository, metricsFactory MetricsFactory) *MetricsService {
	return &MetricsService{metricsRepository: metricsRepository, metricsFactory: metricsFactory}
}
