package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
	"github.com/mrLandyrev/golang-metrics/internal/server/app/transport/rest"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
)

type ServerApp struct {
	router  *gin.Engine
	address string
}

func (app *ServerApp) Run() {
	http.ListenAndServe(app.address, app.router)
}

func NewServerApp(config ServerConfig) *ServerApp {
	// build dependencies
	metricsRepository := metrics.NewMemoryMetricsRepository()
	metricsFactory := metrics.NewMetricsFactory()
	metricsService := service.NewMetricsService(metricsRepository, metricsFactory)

	router := gin.New()

	// register handlers
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", rest.BuildGetAllMetricHandler(metricsService))
	router.POST("/update/:kind/:name/:value", rest.BuildUpdateMetricHandler(metricsService))
	router.GET("/value/:kind/:name", rest.BuildGetMetricHandler(metricsService))

	return &ServerApp{router: router, address: config.Address}
}
