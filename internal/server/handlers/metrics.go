package handlers

import (
	"fmt"
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/factory"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/types"
	"github.com/mrLandyrev/golang-metrics/pkg/router"
)

func GetUpdateMetricHandler(metricsService service.MetricsService) func(c *router.Context) {
	return func(c *router.Context) {
		err := metricsService.AddRecord(c.PathParams["kind"], c.PathParams["name"], c.PathParams["value"])

		switch err {
		case nil:
			break
		case factory.ErrUnknownMetricKind:
			c.Response.WriteHeader(http.StatusNotImplemented)
			return
		case types.ErrIncorrectMetricValue:
			c.Response.WriteHeader(http.StatusBadRequest)
			return
		default:
			c.Response.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func GetGetMetricHandler(metricsService service.MetricsService) func(c *router.Context) {
	return func(c *router.Context) {
		item, err := metricsService.GetRecord(c.PathParams["kind"], c.PathParams["name"])

		if err != nil {
			c.Response.WriteHeader(http.StatusInternalServerError)
			return
		}

		if item == nil {
			c.Response.WriteHeader(http.StatusNotFound)
			return
		}

		fmt.Fprint(c.Response, item.GetValue())
	}
}
