package client

import (
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type HttpClient struct {
	httpClient http.Client
}

func (client *HttpClient) SendMetric(metric models.Metric) error {
	response, err := client.httpClient.Post("http://localhost:8080/update/"+metric.GetKind()+"/"+metric.GetName()+"/"+metric.GetValue(), "text/plain-text", nil)
	response.Body.Close()

	return err
}

func NewHttpClient() *HttpClient {
	return &HttpClient{httpClient: http.Client{}}
}
