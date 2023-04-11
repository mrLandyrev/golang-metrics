package service

import "github.com/mrLandyrev/golang-metrics/internal/agent/metrics/collect/exporers"

type CollectService interface {
	RegisterExporter(exporter exporers.Exporter)
	Collect() error
}
