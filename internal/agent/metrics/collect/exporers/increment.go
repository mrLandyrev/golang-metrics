package exporers

import "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/types"

type IncrementExproter struct {
	counter *types.CounterMetric
}

func (exporter *IncrementExproter) Collect() error {
	exporter.counter.AddValue("1")
	return []types.Metric{exporter.counter}, nil
}

func NewIncrementExproter() *IncrementExproter {
	return &IncrementExproter{counter: types.NewCounterMetric("PollCount")}
}
