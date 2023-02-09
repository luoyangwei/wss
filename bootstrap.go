package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
)

type Bootstrap struct {
	addr string

	handlers []IHandler

	// Is it a reused body
	reusedBody bool
}

type BootstrapOptions struct {
	// Addr eg :8080
	Addr string

	// Is it a reused body
	ReusedBody bool
}

func New(opt *BootstrapOptions) *Bootstrap {
	return &Bootstrap{
		reusedBody: opt.ReusedBody,
	}
}

func (b *Bootstrap) AddHandle(h IHandler) {
	b.handlers = append(b.handlers, h)
}

func (b *Bootstrap) Bind(addr string) {
	http.HandleFunc("/", anyHandler)
	attachmentHandler(b)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Panicln(err)
		return
	}
}

func attachmentHandler(b *Bootstrap) {
	length := len(b.handlers)
	for i := 0; i < length; i++ {
		h := b.handlers[i]
		for pattern, handler := range h.Route() {

			// intercept
			http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
				//go func() {

				h.addResponseWriter(writer)

				// Get from the request body headers
				httpHeader := request.Header
				r := &Request{headers: make(map[string]string, 100)}
				for name, value := range httpHeader {
					var headerBuilder strings.Builder
					for _, v := range value {
						headerBuilder.WriteString(v)
					}
					r.headers[name] = headerBuilder.String()
				}

				// Get the data bytes of body
				if b.reusedBody {
					bytesBody, _ := io.ReadAll(request.Body)
					reader := io.NopCloser(bytes.NewBuffer(bytesBody))
					r.body = reader
				} else {
					r.body = request.Body
				}

				// Call the actual method
				handler(r)

				//}()
			})

		}
	}
}

func anyHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		//time.Sleep(2 * time.Second)
		//count++
		//log.Println("end count: ", count)
	}()
}
