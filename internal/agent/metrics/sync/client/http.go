package client

import (
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type HTTPClient struct {
	httpClient http.Client
}

func (client *HTTPClient) SendMetric(metric models.Metric) error {
	response, err := client.httpClient.Post("http://localhost:8080/update/"+metric.GetKind()+"/"+metric.GetName()+"/"+metric.GetValue(), "text/plain-text", nil)
	response.Body.Close()

	return err
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{httpClient: http.Client{}}
}
