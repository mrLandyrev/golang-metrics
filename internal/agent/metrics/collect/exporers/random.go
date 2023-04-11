package exporers

import (
	"math/rand"
	"strconv"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/types"
)

type RandomExproter struct {
	counter *types.GaugeMetric
}

func (exporter *RandomExproter) GetMetrics() ([]types.Metric, error) {
	exporter.counter.AddValue(strconv.FormatFloat(rand.Float64(), 'f', -1, 64))
	return []types.Metric{exporter.counter}, nil
}

func NewRandomExproter() *RandomExproter {
	return &RandomExproter{counter: types.NewGaugeMetric("RandomValue")}
}
