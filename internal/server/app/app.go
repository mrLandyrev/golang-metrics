package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
	"github.com/mrLandyrev/golang-metrics/internal/metrics/storage"
	"github.com/mrLandyrev/golang-metrics/internal/server/app/transport/rest"
	"github.com/mrLandyrev/golang-metrics/internal/server/metrics/service"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v5/stdlib"
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
		os.Exit(0)
	}

	db, err := sql.Open("pgx", config.DatabaseConnection)
	if err != nil {
		logger.Fatal("Cannot connect database")
	}

	var metricsRepository service.MetricsRepository
	// build dependencies
	if config.DatabaseConnection != "" {
		metricsRepository, _ = storage.NewDatabaseMetricsRepository(db)
	} else if config.FileStoragePath != "" {
		metricsRepository, _ = storage.NewFileMetricsRepository(config.FileStoragePath, config.StoreInterval, config.NeedRestore)
	} else {
		metricsRepository = storage.NewMemoryMetricsRepository()
	}
	metricsFactory := metrics.NewMetricsFactory()
	metricsService := service.NewMetricsService(metricsRepository, metricsFactory)

	router := gin.New()

	router.Use(rest.LoggingMiddleware(logger))
	router.Use(rest.GzipMiddleware())

	// register handlers
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", rest.BuildGetAllMetricHandler(metricsService))
	router.POST("/update/:kind/:name/:value", rest.BuildUpdateMetricHandler(metricsService))
	router.POST("/update", rest.BuildJSONUpdateMetricHandler(metricsService))
	router.POST("/value", rest.BuildJSONGetMetricHandler(metricsService))
	router.POST("/updates", rest.BuildJSONBatchUpdateMetricsHandler(metricsService))
	router.GET("/value/:kind/:name", rest.BuildGetMetricHandler(metricsService))
	router.GET("/ping", rest.BuildPingHandler(db))

	return &ServerApp{router: router, address: config.Address}
}
