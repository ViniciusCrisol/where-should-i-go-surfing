package test

import (
	"bytes"
	"io"
	"net/http"
)

type ResponseBuilder struct {
	response *http.Response
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		response: &http.Response{},
	}
}

func (builder *ResponseBuilder) Build() *http.Response {
	return builder.response
}

func (builder *ResponseBuilder) Status(status int) *ResponseBuilder {
	builder.response.StatusCode = status
	return builder
}

func (builder *ResponseBuilder) StrBody(body string) *ResponseBuilder {
	builder.response.Body = io.NopCloser(bytes.NewBufferString(body))
	return builder
}
