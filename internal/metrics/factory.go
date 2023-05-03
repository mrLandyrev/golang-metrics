package metrics

type MetricsFactory struct{}

func (factory *MetricsFactory) GetInstance(kind string, name string) (Metric, error) {
	switch kind {
	case "counter":
		return NewCounterMetric(name), nil
	case "gauge":
		return NewGaugeMetric(name), nil
	default:
		return nil, ErrUnknownMetricKind
	}
}

func NewMetricsFactory() *MetricsFactory {
	return &MetricsFactory{}
}
