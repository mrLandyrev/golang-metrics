package client

import (
	"fmt"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type FmtClient struct{}

func (client *FmtClient) SendMetric(metric models.Metric) error {
	fmt.Println(metric.GetName(), metric.GetKind(), metric.GetValue())
	return nil
}

func NewFmtClient() *FmtClient {
	return &FmtClient{}
}
