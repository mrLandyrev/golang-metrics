package metrics

type MemoryRepository struct {
	data []Metric
}

func (storage *MemoryRepository) GetByKindAndName(kind string, name string) (Metric, error) {
	for _, item := range storage.data {
		if item.GetKind() == kind && item.GetName() == name {
			return item, nil
		}
	}

	return nil, nil
}

func (storage *MemoryRepository) Persist(item Metric) error {
	for index, storedItem := range storage.data {
		if storedItem.GetKind() == item.GetKind() && storedItem.GetName() == item.GetName() {
			storage.data[index] = item

			return nil
		}
	}
	storage.data = append(storage.data, item)

	return nil
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}
