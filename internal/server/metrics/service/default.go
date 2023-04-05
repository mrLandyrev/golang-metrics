package service

import (
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/factory"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/repository"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/types"
)

type DefaultMetricsService struct {
	metricsRepository repository.MetricsRepository
	metricsFactory    factory.MetricsFactory
}

func (service *DefaultMetricsService) AddRecord(kind string, name string, value string) error {
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

func (service *DefaultMetricsService) GetRecord(kind string, name string) (types.Metric, error) {
	return service.metricsRepository.GetByKindAndName(kind, name)
}

func NewDefaultMetricsService(metricsRepository repository.MetricsRepository, metricsFactory factory.MetricsFactory) *DefaultMetricsService {
	return &DefaultMetricsService{metricsRepository: metricsRepository, metricsFactory: metricsFactory}
}
