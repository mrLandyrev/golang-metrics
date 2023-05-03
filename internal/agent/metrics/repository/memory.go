package repository

import (
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type MemoryMetricsRepository struct {
	data []models.Metric
}

func (storage *MemoryMetricsRepository) GetAll() ([]models.Metric, error) {
	return storage.data, nil
}

func (storage *MemoryMetricsRepository) GetByKindAndName(kind string, name string) (models.Metric, error) {
	for _, item := range storage.data {
		if item.GetKind() == kind && item.GetName() == name {
			return item, nil
		}
	}

	return nil, nil
}

func (storage *MemoryMetricsRepository) Persist(item models.Metric) error {
	for index, storedItem := range storage.data {
		if storedItem.GetKind() == item.GetKind() && storedItem.GetName() == item.GetName() {
			storage.data[index] = item

			return nil
		}
	}
	storage.data = append(storage.data, item)

	return nil
}

func (storage *MemoryMetricsRepository) Clear() error {
	storage.data = []models.Metric{}

	return nil
}

func NewMemoryMetricsRepository() *MemoryMetricsRepository {
	return &MemoryMetricsRepository{}
}
