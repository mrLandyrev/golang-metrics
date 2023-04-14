package exporters

import (
	"runtime"
	"strconv"

	"github.com/mrLandyrev/golang-metrics/internal/agent/metrics/models"
)

type RuntimeExproter struct {
	rtm runtime.MemStats
}

func (exporter *RuntimeExproter) GetMetrics() ([]models.Metric, error) {
	runtime.ReadMemStats(&exporter.rtm)

	metricsMap := map[string]float64{}

	metricsMap["Alloc"] = float64(exporter.rtm.Alloc)
	metricsMap["BuckHashSys"] = float64(exporter.rtm.BuckHashSys)
	metricsMap["Frees"] = float64(exporter.rtm.Frees)
	metricsMap["GCCPUFraction"] = float64(exporter.rtm.GCCPUFraction)
	metricsMap["GCSys"] = float64(exporter.rtm.GCSys)
	metricsMap["HeapAlloc"] = float64(exporter.rtm.HeapAlloc)
	metricsMap["HeapIdle"] = float64(exporter.rtm.HeapIdle)
	metricsMap["HeapInuse"] = float64(exporter.rtm.HeapInuse)
	metricsMap["HeapObjects"] = float64(exporter.rtm.HeapObjects)
	metricsMap["HeapReleased"] = float64(exporter.rtm.HeapReleased)
	metricsMap["HeapSys"] = float64(exporter.rtm.HeapSys)
	metricsMap["LastGC"] = float64(exporter.rtm.LastGC)
	metricsMap["Lookups"] = float64(exporter.rtm.Lookups)
	metricsMap["MCacheInuse"] = float64(exporter.rtm.MCacheInuse)
	metricsMap["MCacheSys"] = float64(exporter.rtm.MCacheSys)
	metricsMap["MSpanInuse"] = float64(exporter.rtm.MSpanInuse)
	metricsMap["MSpanSys"] = float64(exporter.rtm.MSpanSys)
	metricsMap["Mallocs"] = float64(exporter.rtm.Mallocs)
	metricsMap["NextGC"] = float64(exporter.rtm.NextGC)
	metricsMap["NumForcedGC"] = float64(exporter.rtm.NumForcedGC)
	metricsMap["NumGC"] = float64(exporter.rtm.NumGC)
	metricsMap["OtherSys"] = float64(exporter.rtm.OtherSys)
	metricsMap["PauseTotalNs"] = float64(exporter.rtm.PauseTotalNs)
	metricsMap["StackInuse"] = float64(exporter.rtm.StackInuse)
	metricsMap["StackSys"] = float64(exporter.rtm.StackSys)
	metricsMap["Sys"] = float64(exporter.rtm.Sys)
	metricsMap["TotalAlloc"] = float64(exporter.rtm.TotalAlloc)

	res := []models.Metric{}

	for key, value := range metricsMap {
		metric := models.NewGaugeMetric(key)
		metric.AddValue(strconv.FormatFloat(value, 'f', -1, 64))
		res = append(res, metric)
	}

	return res, nil
}

func NewRuntimeExporter() *RuntimeExproter {
	return &RuntimeExproter{}
}
