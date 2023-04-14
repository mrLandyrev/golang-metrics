package client

import (
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type HTTPClient struct {
	httpClient http.Client
	addr       string
}

func (client *HTTPClient) SendMetric(metric models.Metric) error {
	response, err := client.httpClient.Post("http://"+client.addr+"/update/"+metric.GetKind()+"/"+metric.GetName()+"/"+metric.GetValue(), "text/plain-text", nil)
	response.Body.Close()

	return err
}

func NewHTTPClient(addr string) *HTTPClient {
	return &HTTPClient{httpClient: http.Client{}, addr: addr}
}
