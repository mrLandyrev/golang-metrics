package metrics

type Service struct {
	repository Repository
	factory    *Factory
}

func (service *Service) AddRecord(kind string, name string, value string) error {
	item, err := service.repository.GetByKindAndName(kind, name)

	if err != nil {
		return err
	}

	if item == nil {
		item, err = (*service.factory).GetMetric(kind, name)

		if err != nil {
			return err
		}
	}

	err = item.AddStrValue(value)

	if err != nil {
		return err
	}

	return service.repository.Persist(item)
}

func (service *Service) GetRecord(kind string, name string) (Metric, error) {
	return service.repository.GetByKindAndName(kind, name)
}

func NewService(repository Repository, factory *Factory) *Service {
	return &Service{repository: repository, factory: factory}
}
