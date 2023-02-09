package wss

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

var (
	_ Requester = &Request{}
)

type Handlers map[string]string

type Requester interface {
	// GetHeaders 获取头信息
	GetHeaders() map[string]string
	GetHeaderByName(name string) string

	// GetBody 获取body信息
	GetBody() []byte

	// GetParameters 获取请求参数
	GetParameters() map[string]string
	GetParameterByName(name string) string
}

type Request struct {

	// originRequest Original request
	originRequest *http.Request
	body          io.ReadCloser

	headers map[string]string
}

// setHeaders Get from the request body headers
func (r *Request) setHeaders() {
	httpHeader := r.originRequest.Header
	for name, value := range httpHeader {
		var headerBuilder strings.Builder
		for _, v := range value {
			headerBuilder.WriteString(v)
		}
		r.headers[name] = headerBuilder.String()
	}
}

// setRequestBody Get the data bytes of body
func (r *Request) setRequestBody(reused bool) {
	if reused {
		bytesBody, _ := io.ReadAll(r.originRequest.Body)
		reader := io.NopCloser(bytes.NewBuffer(bytesBody))
		r.body = reader
	} else {
		r.body = r.originRequest.Body
	}
}

func (r *Request) GetHeaders() map[string]string {
	return r.headers
}

func (r *Request) GetHeaderByName(name string) string {
	return r.headers[name]
}

func (r *Request) GetBody() []byte {
	bytesBody, _ := io.ReadAll(r.body)
	return bytesBody
}

func (r *Request) GetParameters() map[string]string {
	return nil
}

func (r *Request) GetParameterByName(name string) string {
	return ""
}
