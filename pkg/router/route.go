package router

import (
	"strings"
)

type Route struct {
	method   string
	segments []string
	handler  func(c *Context)
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

func (route *Route) Handle(c *Context) {
	route.handler(c)
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
