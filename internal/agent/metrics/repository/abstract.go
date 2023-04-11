package repository

import "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/types"

type MetricsRepository interface {
	GetAll() ([]types.Metric, error)
	GetByKindAndName(kind string, name string) (types.Metric, error)
	Persist(item types.Metric) error
	Clear() error
}
