package rest

import (
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type HTTPClient struct {
	httpClient http.Client
	addr       string
}

func (client *HTTPClient) SendMetric(metric metrics.Metric) error {
	response, err := client.httpClient.Post("http://"+client.addr+"/update/"+metric.Kind()+"/"+metric.Name()+"/"+metric.Value(), "plain/text", nil)
	response.Body.Close()

	return err
}

func NewHTTPClient(addr string) *HTTPClient {
	return &HTTPClient{httpClient: http.Client{}, addr: addr}
}
