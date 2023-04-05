package factory

import "github.com/mrLandyrev/golang-metrics/internal/server/metrics/types"

type DefaultMetricsFactory struct{}

func (factory *DefaultMetricsFactory) GetInstance(kind string, name string) (types.Metric, error) {
	switch kind {
	case "counter":
		return types.NewCounterMetric(name), nil
	case "gauge":
		return types.NewGaugeMetric(name), nil
	default:
		return nil, ErrUnknownMetricKind
	}
}

func NewDefaultMetricsFactory() *DefaultMetricsFactory {
	return &DefaultMetricsFactory{}
}
