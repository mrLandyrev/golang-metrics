package repository

import "github.com/mrLandyrev/golang-metrics/internal/server/metrics/types"

type MetricsRepository interface {
	GetByKindAndName(kind string, name string) (types.Metric, error)
	Persist(item types.Metric) error
}
