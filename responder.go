package wss

import (
	"net/http"
)

// Action 目的
type Action func(r *Request)

// Responder This is the responder
type Responder interface {

	// Route Set route
	Route() Routers

	// Response response content
	Response(resp []byte)

	//decomposeResponder()
	addResponseWriter(w http.ResponseWriter)
}

// ResponseEmpowerment Let Action operate some capabilities of Response
type ResponseEmpowerment struct {
	writer http.ResponseWriter
}

func initResponder(resp Responder, boot *Bootstrap) {
	routes := resp.Route()
	for _, route := range routes {

		// intercept
		http.HandleFunc(route.Pattern, func(writer http.ResponseWriter, request *http.Request) {
			req := &Request{headers: make(map[string]string, 100), originRequest: request}
			resp.addResponseWriter(writer)

			// Parse data
			req.setHeaders()
			req.setRequestBody(boot.reusedBody)

			// Call the actual method
			route.Action(req)
		})
	}
}

func (resp *ResponseEmpowerment) addResponseWriter(w http.ResponseWriter) {
	resp.writer = w
}

func (resp *ResponseEmpowerment) Response(buf []byte) {
	w := resp.writer
	_, _ = w.Write(buf)
	return
}
