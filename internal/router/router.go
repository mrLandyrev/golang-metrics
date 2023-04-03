package router

import (
	"net/http"
	"regexp"
	"strings"
)

type Context struct {
	method     string
	pathParams map[string]string
}

func (context *Context) GetPathParam(name string) string {
	return context.pathParams[name]
}

type Route struct {
	method   string
	segments []string
	handler  func(w http.ResponseWriter, r *http.Request, c *Context)
}

func (route *Route) Compare(method string, segments []string) bool {
	if method != route.method {
		return false
	}

	if len(segments) != len(route.segments) {
		return false
	}

	for index := range segments {
		if strings.HasPrefix(route.segments[index], ":") {
			continue
		}

		if route.segments[index] == segments[index] {
			continue
		}

		return false
	}

	return true
}

func (route *Route) CalculatePathParams(segments []string) map[string]string {
	res := make(map[string]string)

	pathParamName := ""
	pathParamValue := ""

	for index := range segments {
		if strings.HasPrefix(route.segments[index], ":") {
			pathParamName = strings.TrimPrefix(route.segments[index], ":")
			pathParamValue = segments[index]

			res[pathParamName] = pathParamValue
		}
	}

	return res
}

type Router struct {
	routes []*Route
}

func (router *Router) Use(method string, path string, handler func(w http.ResponseWriter, r *http.Request, c *Context)) {
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
			context := &Context{method: method, pathParams: route.CalculatePathParams(segments)}
			route.handler(w, r, context)
			return
		}
	}

	http.NotFound(w, r)
}

func NewRouter() *Router {
	return &Router{}
}
