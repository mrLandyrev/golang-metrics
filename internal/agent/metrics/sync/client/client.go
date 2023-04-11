package client

import "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/types"

type Client interface {
	SendMetric(metric types.Metric) error
}
