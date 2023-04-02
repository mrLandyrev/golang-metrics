package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var (
	couter = map[string]int64{}
	guage  = map[string]float64{}
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/update/", handleUpdate)
	router.HandleFunc("/get/", handleGet)

	http.ListenAndServe(":8080", router)
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	metricType := segments[2]
	metricName := segments[3]

	switch metricType {
	case "counter":
		metricValue, _ := strconv.ParseInt(segments[4], 10, 64)
		couter[metricName] += metricValue
	case "guage":
		metricValue, _ := strconv.ParseFloat(segments[4], 64)
		guage[metricName] = metricValue
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	metricType := segments[2]
	metricName := segments[3]

	switch metricType {
	case "counter":
		fmt.Fprint(w, couter[metricName])
	case "guage":
		fmt.Fprint(w, guage[metricName])
	}
}
