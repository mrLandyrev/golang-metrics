package rest

import (
	"strconv"

	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type Metric struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func From(record metrics.Metric) Metric {
	m := Metric{}

	m.ID = record.Name()
	m.MType = record.Kind()

	switch record.Kind() {
	case "counter":
		delta, _ := strconv.ParseInt(record.Value(), 10, 64)
		m.Delta = &delta
	case "gauge":
		gauge, _ := strconv.ParseFloat(record.Value(), 64)
		m.Value = &gauge
	}

	return m
}

func (m *Metric) To() (metrics.Metric, error) {
	switch m.MType {
	case "counter":
		return &metrics.CounterMetric{NameData: m.ID, ValueData: *m.Delta}, nil
	case "gauge":
		return &metrics.GaugeMetric{NameData: m.ID, ValueData: *m.Value}, nil
	default:
		return nil, metrics.ErrUnknownMetricKind
	}
}
