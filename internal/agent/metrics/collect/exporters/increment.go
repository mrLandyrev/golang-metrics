package exporters

import (
	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type IncrementExproter struct{}

func (exporter *IncrementExproter) GetMetrics() ([]models.Metric, error) {
	metric := models.NewCounterMetric("PollCount")
	err := metric.AddValue("1")
	if err != nil {
		return nil, err
	}

	return []models.Metric{metric}, nil
}

func NewIncrementExproter() *IncrementExproter {
	return &IncrementExproter{}
}
