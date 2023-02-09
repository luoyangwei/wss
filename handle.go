package main

import (
	"net/http"
)

type H func(r *Request)

type IHandler interface {
	Route() map[string]H

	// Response 响应内容
	Response(resp []byte)

	addResponseWriter(w http.ResponseWriter)
}

type Handle struct {
	writer http.ResponseWriter
}

func (h *Handle) addResponseWriter(w http.ResponseWriter) {
	h.writer = w
}

func (h *Handle) Response(resp []byte) {
	w := h.writer
	_, _ = w.Write(resp)
	return
}
