package metrics

type MemoryMetricsRepository struct {
	data []Metric
}

func (storage *MemoryMetricsRepository) GetAll() ([]Metric, error) {
	return storage.data, nil
}

func (storage *MemoryMetricsRepository) GetByKindAndName(kind string, name string) (Metric, error) {
	for _, item := range storage.data {
		if item.Kind() == kind && item.Name() == name {
			return item, nil
		}
	}

	return nil, nil
}

func (storage *MemoryMetricsRepository) Persist(item Metric) error {
	for index, storedItem := range storage.data {
		if storedItem.Kind() == item.Kind() && storedItem.Name() == item.Name() {
			storage.data[index] = item

			return nil
		}
	}
	storage.data = append(storage.data, item)

	return nil
}

func (storage *MemoryMetricsRepository) Clear() error {
	storage.data = []Metric{}

	return nil
}

func NewMemoryMetricsRepository() *MemoryMetricsRepository {
	return &MemoryMetricsRepository{}
}
