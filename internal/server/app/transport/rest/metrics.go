package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

// -- dependencies --

type MetricsService interface {
	GetAll() ([]metrics.Metric, error)
	AddRecord(kind string, name string, value string) error
	GetRecord(kind string, name string) (metrics.Metric, error)
}

// -- dependencies --

func BuildUpdateMetricHandler(metricsService MetricsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := metricsService.AddRecord(c.Param("kind"), c.Param("name"), c.Param("value"))

		switch err {
		case nil:
			break
		case metrics.ErrUnknownMetricKind:
			c.Status(http.StatusNotImplemented)
			return
		case metrics.ErrIncorrectMetricValue:
			c.Status(http.StatusBadRequest)
			return
		default:
			c.Status(http.StatusInternalServerError)
			return
		}
	}
}

func BuildGetMetricHandler(metricsService MetricsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		item, err := metricsService.GetRecord(c.Param("kind"), c.Param("name"))

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if item == nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.String(http.StatusOK, item.Value())
	}
}

func BuildGetAllMetricHandler(MetricsService MetricsService) gin.HandlerFunc {
	return func(c *gin.Context) {

		items, err := MetricsService.GetAll()

		if err != nil {
			c.Status(http.StatusInternalServerError)
		}

		metrics := make([]struct {
			Name  string
			Kind  string
			Value string
		}, 0, len(items))

		for _, item := range items {
			metrics = append(metrics, struct {
				Name  string
				Kind  string
				Value string
			}{
				Name:  item.Name(),
				Kind:  item.Kind(),
				Value: item.Value(),
			})
		}

		c.HTML(http.StatusOK, "list.html", gin.H{
			"Metrics": metrics,
		})
	}
}
