package exporers

import "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/types"

type RuntimeExproter struct{}

func (exporter *RuntimeExproter) GetMetrics() ([]types.Metric, error) {
	test := types.NewCounterMetric("test")
	test.AddValue("4")
	return []types.Metric{
		test,
	}, nil
}

func NewRuntimeExporter() *RuntimeExproter {
	return &RuntimeExproter{}
}
