package metrics

import "errors"

var (
	ErrUnknownMetricKind = errors.New("unknown metric kind detected")
)

type Factory struct{}

func (factory *Factory) GetMetric(kind string, name string) (Metric, error) {
	switch kind {
	case "counter":
		return NewCounterMetric(name), nil
	case "gauge":
		return NewGaugeMetric(name), nil
	default:
		return nil, ErrUnknownMetricKind
	}
}

func NewFactory() *Factory {
	return &Factory{}
}