package rest

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type HTTPClient struct {
	httpClient http.Client
	addr       string
}

func (client *HTTPClient) SendMetric(metric metrics.Metric) error {
	var body bytes.Buffer
	jBody, err := json.Marshal(From(metric))
	if err != nil {
		return err
	}
	gzipWriter := gzip.NewWriter(&body)
	gzipWriter.Write(jBody)
	gzipWriter.Flush()
	req, err := http.NewRequest(http.MethodPost, "http://"+client.addr+"/update/", &body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Encoding", "gzip")
	response, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	response.Body.Close()

	return err
}

func NewHTTPClient(addr string) *HTTPClient {
	return &HTTPClient{httpClient: http.Client{}, addr: addr}
}
