package router

import "net/http"

type Context struct {
	PathParams map[string]string
	Request    *http.Request
	Response   http.ResponseWriter
}

func NewContext(pathParams map[string]string, request *http.Request, response http.ResponseWriter) *Context {
	return &Context{PathParams: pathParams, Request: request, Response: response}
}
