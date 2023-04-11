package factory

import (
	"errors"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/types"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/models"
)

var (
	ErrUnknownMetricKind    = errors.New("unknow metric kind")
	ErrIncorrectMetricValue = errors.New("incorrect metric value")
)

type MetricsFactory struct{}

func (factory *MetricsFactory) GetInstance(kind string, name string) (models.Metric, error) {
	switch kind {
	case "counter":
		return types.NewCounterMetric(name), nil
	case "gauge":
		return types.NewGaugeMetric(name), nil
	default:
		return nil, ErrUnknownMetricKind
	}
}

func NewMetricsFactory() *MetricsFactory {
	return &MetricsFactory{}
}
