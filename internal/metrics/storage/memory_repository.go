package storage

import (
	"sync"

	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type MemoryMetricsRepository struct {
	data  []metrics.Metric
	mutex sync.Mutex
}

func (storage *MemoryMetricsRepository) GetAll() ([]metrics.Metric, error) {
	storage.mutex.Lock()
	data := storage.data
	storage.mutex.Unlock()

	return data, nil
}

func (storage *MemoryMetricsRepository) GetByKindAndName(kind string, name string) (metrics.Metric, error) {
	storage.mutex.Lock()
	for _, item := range storage.data {
		if item.Kind() == kind && item.Name() == name {
			storage.mutex.Unlock()
			return item, nil
		}
	}

	storage.mutex.Unlock()
	return nil, nil
}

func (storage *MemoryMetricsRepository) Persist(item metrics.Metric) error {
	storage.mutex.Lock()
	for index, storedItem := range storage.data {
		if storedItem.Kind() == item.Kind() && storedItem.Name() == item.Name() {
			storage.data[index] = item
			storage.mutex.Unlock()

			return nil
		}
	}
	storage.data = append(storage.data, item)
	storage.mutex.Unlock()

	return nil
}

func (storage *MemoryMetricsRepository) Clear() error {
	storage.mutex.Lock()
	storage.data = []metrics.Metric{}
	storage.mutex.Unlock()

	return nil
}

func NewMemoryMetricsRepository() *MemoryMetricsRepository {
	return &MemoryMetricsRepository{}
}
