package exporters

import (
	"math/rand"
	"strconv"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type RandomExproter struct{}

func (exporter *RandomExproter) GetMetrics() ([]models.Metric, error) {
	metric := models.NewGaugeMetric("RandomValue")
	metric.AddValue(strconv.FormatFloat(rand.Float64(), 'f', -1, 64))

	return []models.Metric{metric}, nil
}

func NewRandomExproter() *RandomExproter {
	return &RandomExproter{}
}
