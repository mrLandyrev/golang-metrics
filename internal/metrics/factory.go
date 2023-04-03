package metrics

import "errors"

type Factory struct{}

func (factory *Factory) GetMetric(kind string, name string) (Metric, error) {
	switch kind {
	case "counter":
		return NewCounterMetric(name), nil
	case "guage":
		return NewGuageMetric(name), nil
	default:
		return nil, errors.New("NOT IMPLEMENTED")
	}
}

func NewFactory() *Factory {
	return &Factory{}
}
