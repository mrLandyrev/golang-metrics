package client

import (
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type HttpClient struct{}

func (client *HttpClient) SendMetric(metric models.Metric) error {
	_, err := http.Post("http://localhost:8080/update/"+metric.GetKind()+"/"+metric.GetName()+"/"+metric.GetValue(), "text/plain-text", nil)
	return err
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}
