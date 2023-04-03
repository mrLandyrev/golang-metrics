package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

var (
	metricsService *metrics.Service
)

func main() {
	metricsService = metrics.NewService(metrics.NewMemoryRepository(), metrics.NewFactory())

	router := http.NewServeMux()
	router.HandleFunc("/update/", handleUpdate)
	router.HandleFunc("/get/", handleGet)

	http.ListenAndServe(":8080", router)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	kind := segments[2]
	name := segments[3]
	value := segments[4]

	err := metricsService.AddRecord(kind, name, value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	kind := segments[2]
	name := segments[3]

	item, err := metricsService.GetRecord(kind, name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Metric not found")
		return
	}

	fmt.Fprint(w, item.GetStrValue())
}
