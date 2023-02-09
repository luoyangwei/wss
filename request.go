package main

import "io"

var (
	_ IRequester = &Request{}
)

type IRequester interface {
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
	headers map[string]string
	body    io.ReadCloser
}

func (r *Request) GetParameterByName(name string) string {
	return ""
}

func (r *Request) GetHeaders() map[string]string {
	return nil
}

func (r *Request) GetHeaderByName(name string) string {
	return ""
}

func (r *Request) GetBody() []byte {
	return nil
}
func (r *Request) GetParameters() map[string]string {
	return nil
}

func (r *Request) funcGetParameterByName(name string) string {
	return ""
}
