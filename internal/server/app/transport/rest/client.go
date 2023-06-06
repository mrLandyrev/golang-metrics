package rest

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"net/http"

	retry "github.com/mrLandyrev/golang-metrics/internal"
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

func (client *HTTPClient) SendMetrics(metrics []metrics.Metric) error {
	var bodyData []Metric
	for _, metric := range metrics {
		bodyData = append(bodyData, From(metric))
	}
	jBody, err := json.Marshal(bodyData)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	gzipWriter := gzip.NewWriter(&body)
	gzipWriter.Write(jBody)
	gzipWriter.Flush()
	req, err := http.NewRequest(http.MethodPost, "http://"+client.addr+"/updates/", &body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Encoding", "gzip")
	var response *http.Response
	err = retry.HandleFunc(func() error {
		var err error
		response, err = client.httpClient.Do(req)
		defer response.Body.Close()
		if err != nil {
			return err
		}

		return err
	}, 4, nil)

	if err != nil {
		return err
	}

	return err
}

func NewHTTPClient(addr string) *HTTPClient {
	return &HTTPClient{httpClient: http.Client{}, addr: addr}
}
