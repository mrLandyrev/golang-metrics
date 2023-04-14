package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/factory"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/models"
)

type MetricsService interface {
	GetAll() ([]models.Metric, error)
	AddRecord(kind string, name string, value string) error
	GetRecord(kind string, name string) (models.Metric, error)
}

func BuildUpdateMetricHandler(metricsService MetricsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := metricsService.AddRecord(c.Param("kind"), c.Param("name"), c.Param("value"))

		switch err {
		case nil:
			break
		case factory.ErrUnknownMetricKind:
			c.Status(http.StatusNotImplemented)
			return
		case factory.ErrIncorrectMetricValue:
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

		c.String(http.StatusOK, item.GetValue())
	}
}

func BuildGetAllMetricHandler(MetricsService MetricsService) gin.HandlerFunc {
	return func(c *gin.Context) {

		items, err := MetricsService.GetAll()

		if err != nil {
			c.Status(http.StatusInternalServerError)
		}

		metrics := []struct {
			Name  string
			Kind  string
			Value string
		}{}

		for _, item := range items {
			metrics = append(metrics, struct {
				Name  string
				Kind  string
				Value string
			}{
				Name:  item.GetName(),
				Kind:  item.GetKind(),
				Value: item.GetValue(),
			})
		}

		c.HTML(http.StatusOK, "list.html", gin.H{
			"Metrics": metrics,
		})
	}
}
