package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
	"github.com/mrLandyrev/golang-metrics/internal/server/app/transport/rest"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
	"go.uber.org/zap"
)

type ServerApp struct {
	router  *gin.Engine
	address string
}

func (app *ServerApp) Run() {
	http.ListenAndServe(app.address, app.router)
}

func NewServerApp(config ServerConfig) *ServerApp {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("Cannot create logger")
	}
	// build dependencies
	metricsRepository := metrics.NewMemoryMetricsRepository()
	metricsFactory := metrics.NewMetricsFactory()
	metricsService := service.NewMetricsService(metricsRepository, metricsFactory)

	router := gin.New()

	router.Use(rest.BuildLoggingMiddleware(logger))

	// register handlers
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", rest.BuildGetAllMetricHandler(metricsService))
	router.POST("/update/:kind/:name/:value", rest.BuildUpdateMetricHandler(metricsService))
	router.POST("/update", rest.BuildJSONUpdateMetricHandler(metricsService))
	router.POST("/value", rest.BuildJSONGetMetricHandler(metricsService))
	router.GET("/value/:kind/:name", rest.BuildGetMetricHandler(metricsService))

	return &ServerApp{router: router, address: config.Address}
}
