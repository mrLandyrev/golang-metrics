package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

// -- dependencies --

type MetricsService interface {
	GetAll() ([]metrics.Metric, error)
	AddRecord(kind string, name string, value string) (metrics.Metric, error)
	GetRecord(kind string, name string) (metrics.Metric, error)
}

// -- dependencies --

func BuildUpdateMetricHandler(metricsService MetricsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := metricsService.AddRecord(c.Param("kind"), c.Param("name"), c.Param("value"))

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

func BuildJSONUpdateMetricHandler(metricsService MetricsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var m Metric
		err := c.BindJSON(&m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		var recordValue string
		switch m.MType {
		case "gauge":
			recordValue = strconv.FormatFloat(*m.Value, 'f', -1, 64)
		case "counter":
			recordValue = strconv.FormatInt(*m.Delta, 10)
		default:
			c.Status(http.StatusNotImplemented)
			return
		}

		metricValue, err := metricsService.AddRecord(m.MType, m.ID, recordValue)

		switch err {
		case nil:
			c.JSON(http.StatusOK, From(metricValue))
			return
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

func BuildJSONGetMetricHandler(metricsService MetricsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var m Metric
		err := c.BindJSON(&m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		item, err := metricsService.GetRecord(m.MType, m.ID)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if item == nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, From(item))
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
