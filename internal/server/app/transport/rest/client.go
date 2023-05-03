package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type HTTPClient struct {
	httpClient http.Client
	addr       string
}

func (client *HTTPClient) SendMetric(metric metrics.Metric) error {
	body, err := json.Marshal(metric)
	if err != nil {
		return err
	}
	response, err := client.httpClient.Post("http://"+client.addr+"/update/", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	response.Body.Close()

	return err
}

func NewHTTPClient(addr string) *HTTPClient {
	return &HTTPClient{httpClient: http.Client{}, addr: addr}
}
