package rest

import (
	"fmt"
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/factory"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/models"
	"github.com/mrLandyrev/golang-metrics/pkg/router"
)

type MetricsService interface {
	AddRecord(kind string, name string, value string) error
	GetRecord(kind string, name string) (models.Metric, error)
}

func BuildUpdateMetricHandler(metricsService MetricsService) func(c *router.Context) {
	return func(c *router.Context) {
		err := metricsService.AddRecord(c.PathParams["kind"], c.PathParams["name"], c.PathParams["value"])

		switch err {
		case nil:
			break
		case factory.ErrUnknownMetricKind:
			c.Response.WriteHeader(http.StatusNotImplemented)
			return
		case factory.ErrIncorrectMetricValue:
			c.Response.WriteHeader(http.StatusBadRequest)
			return
		default:
			c.Response.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func BuildGetMetricHandler(metricsService MetricsService) func(c *router.Context) {
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
