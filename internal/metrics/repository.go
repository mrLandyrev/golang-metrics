package metrics

type Repository interface {
	GetByKindAndName(kind string, name string) (Metric, error)
	Persist(item Metric) error
}
