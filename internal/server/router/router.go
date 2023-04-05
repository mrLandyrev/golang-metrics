package router

import (
	"net/http"
	"regexp"
	"strings"
)

type Router struct {
	routes []*Route
}

func (router *Router) Use(method string, path string, handler func(c *Context)) {
	if isValid, _ := regexp.MatchString(`^(\/:?[a-zA-Z,\-,_,0-9]+)*$`, path); !isValid {
		panic("Route path invalid")
	}

	segments := strings.Split(path, "/")

	for _, route := range router.routes {
		if route.Compare(method, segments) {
			panic("Routes conflict")
		}
	}

	router.routes = append(router.routes, &Route{method: method, segments: segments, handler: handler})
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	segments := strings.Split(r.URL.Path, "/")

	for _, route := range router.routes {
		if route.Compare(method, segments) {
			context := NewContext(route.CalculatePathParams(segments), r, w)
			route.Handle(context)
			return
		}
	}

	http.NotFound(w, r)
}

func NewRouter() *Router {
	return &Router{}
}
