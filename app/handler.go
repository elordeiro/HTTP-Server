package main

import "slices"

func (server *Server) Handle(req *Request) *Response {
	switch req.Method {
	case MethodGet:
		return server.Get(req)
	// case MethodPost:
	// 	return req.Post()
	// case MethodPut:
	// 	return req.Put()
	// case MethodDelete:
	// 	return req.Delete()
	// case MethodHead:
	// 	return req.Head()
	default:
		return req.NotFound()
	}
}

func (server *Server) Get(req *Request) *Response {
	if slices.Contains(server.Paths, req.Path) {
		return &Response{
			Version: req.Version,
			Status:  StatusOk,
			Reason:  "OK",
			Headers: map[string]string{},
			Body:    "",
		}
	}
	return req.NotFound()
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
