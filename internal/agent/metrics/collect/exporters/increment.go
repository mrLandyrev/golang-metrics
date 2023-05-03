package exporters

import (
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type IncrementExproter struct{}

func (exporter *IncrementExproter) GetMetrics() ([]metrics.Metric, error) {
	metric := metrics.NewCounterMetric("PollCount")
	err := metric.AddValue("1")
	if err != nil {
		return nil, err
	}

	return []metrics.Metric{metric}, nil
}

func NewIncrementExproter() *IncrementExproter {
	return &IncrementExproter{}
}
