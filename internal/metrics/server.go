package metrics

import (
	"fmt"
	"net/http"

	"github.com/mrLandyrev/golang-metrics/internal/server/router"
)

type Server struct {
	metricsService *Service
}

func (server *Server) Listen() {
	router := router.NewRouter()
	router.Use("POST", "/update/:kind/:name/:value", server.updateHandler)
	router.Use("GET", "/get/:kind/:name", server.getHandler)

	http.ListenAndServe(":8080", router)
}

func (server *Server) updateHandler(c *router.Context) {
	err := server.metricsService.AddRecord(c.PathParams["kind"], c.PathParams["name"], c.PathParams["value"])

	switch err {
	case nil:
		break
	case ErrUnknownMetricKind:
		c.Response.WriteHeader(http.StatusNotImplemented)
		return
	case ErrIncorrectMetricValue:
		c.Response.WriteHeader(http.StatusBadRequest)
		return
	default:
		c.Response.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (server *Server) getHandler(c *router.Context) {
	item, err := server.metricsService.GetRecord(c.PathParams["kind"], c.PathParams["name"])

	if err != nil {
		c.Response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if item == nil {
		c.Response.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(c.Response, item.GetStrValue())
}

func NewServer(metricsService *Service) *Server {
	return &Server{metricsService: metricsService}
}
