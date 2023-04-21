package exporters

import (
	"math/rand"
	"strconv"

	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type RandomExproter struct{}

func (exporter *RandomExproter) GetMetrics() ([]metrics.Metric, error) {
	metric := metrics.NewGaugeMetric("RandomValue")
	metric.AddValue(strconv.FormatFloat(rand.Float64(), 'f', -1, 64))

	return []metrics.Metric{metric}, nil
}

func NewRandomExproter() *RandomExproter {
	return &RandomExproter{}
}
