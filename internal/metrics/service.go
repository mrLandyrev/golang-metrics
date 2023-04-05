package metrics

type MetricsService interface {
	AddMetric(kind string, name string, value string) error
	GetMetric(kind string, name string) (Metric, error)
}

type Service struct {
	repository Repository
	factory    *Factory
}

func (service *Service) AddRecord(kind string, name string, value string) error {
	// find item in storage
	item, err := service.repository.GetByKindAndName(kind, name)

	if err != nil {
		return err
	}

	// create item if not found in storage
	if item == nil {
		item, err = (*service.factory).GetMetric(kind, name)

		if err != nil {
			return err
		}
	}

	// apply raw value (how to interpret value encapsulated in metric)
	err = item.AddStrValue(value)

	if err != nil {
		return err
	}

	// save to storage
	return service.repository.Persist(item)
}

func (service *Service) GetRecord(kind string, name string) (Metric, error) {
	return service.repository.GetByKindAndName(kind, name)
}

func NewService(repository Repository, factory *Factory) *Service {
	return &Service{repository: repository, factory: factory}
}
