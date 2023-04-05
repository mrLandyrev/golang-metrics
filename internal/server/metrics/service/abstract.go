package service

import "github.com/mrLandyrev/golang-metrics/internal/server/metrics/types"

type MetricsService interface {
	AddRecord(kind string, name string, value string) error
	GetRecord(kind string, name string) (types.Metric, error)
}
