package repository

import "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/types"

type MemoryMetricsRepository struct {
	data []types.Metric
}

func (storage *MemoryMetricsRepository) GetAll() ([]types.Metric, error) {
	return storage.data, nil
}

func (storage *MemoryMetricsRepository) GetByKindAndName(kind string, name string) (types.Metric, error) {
	for _, item := range storage.data {
		if item.GetKind() == kind && item.GetName() == name {
			return item, nil
		}
	}

	return nil, nil
}

func (storage *MemoryMetricsRepository) Persist(item types.Metric) error {
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
	storage.data = []types.Metric{}

	return nil
}

func NewMemoryMetricsRepository() *MemoryMetricsRepository {
	return &MemoryMetricsRepository{}
}
