package main

import (
	"strconv"
	"strings"
)

func (server *Server) Handle(req *Request) *Response {
	switch req.Method {
	case MethodGet:
		return server.Get(req)
	default:
		return req.NotFound()
	}
}

func (server *Server) Get(req *Request) *Response {
	if req.Path == "/" {
		if handler, ok := server.Paths["/"]; ok {
			return handler(req)
		} else {
			return req.NotFound()
		}
	}

	parts := strings.Split(req.Path, "/")
	if handler, ok := server.Paths["/"+parts[1]]; ok {
		return handler(req)
	}
	return req.NotFound()
}

func (server *Server) AddPath(path string, handler func(*Request) *Response) {
	server.Paths[path] = handler
}

func (req *Request) Ok() *Response {
	return &Response{
		Version: req.Version,
		Status:  StatusOk,
		Reason:  "OK",
		Headers: map[string]string{},
		Body:    "",
	}
}

func (req *Request) Echo() *Response {
	headers := map[string]string{}
	headers["Content-Type"] = "text/plain"
	body := strings.TrimPrefix(req.Path, "/echo/")
	headers["Content-Length"] = strconv.Itoa(len(body))
	return &Response{
		Version: req.Version,
		Status:  StatusOk,
		Reason:  "OK",
		Headers: headers,
		Body:    body,
	}
}

func (req *Request) UserAgent() *Response {
	headers := map[string]string{}
	headers["Content-Type"] = "text/plain"
	body := req.Headers["user-agent:"]
	headers["Content-Length"] = strconv.Itoa(len(body))
	return &Response{
		Version: req.Version,
		Status:  StatusOk,
		Reason:  "OK",
		Headers: headers,
		Body:    body,
	}
}

func (req *Request) NotFound() *Response {
	return &Response{
		Version: req.Version,
		Status:  StatusNotFound,
		Reason:  "Not Found",
		Headers: map[string]string{},
		Body:    "",
	}
}
