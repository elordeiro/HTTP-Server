package main

import (
	"bufio"
	"io"
	"net"
)

// Constants ------------------------------------------------------------------
const (
	MethodGet    = "GET"
	MethodPut    = "PUT"
	MethodPost   = "POST"
	MethodHead   = "HEAD"
	MethodDelete = "DELETE"
)

const (
	StatusOk         = 200
	StatusBadRequest = 400
	StatusNotFound   = 404
)

// ----------------------------------------------------------------------------

// Custom Types ---------------------------------------------------------------
type Server struct {
	Listener net.Listener
	Paths    []string
}

type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    string
}

type Response struct {
	Version string
	Status  int
	Reason  string
	Headers map[string]string
	Body    string
}

// ----------------------------------------------------------------------------

// Reader and Writer ----------------------------------------------------------
type Reader struct {
	reader *bufio.Reader
}

type Writer struct {
	writer io.Writer
}

func NewReader(rd io.Reader) *Reader {
	return &Reader{
		reader: bufio.NewReader(rd),
	}
}

func NewWriter(wt io.Writer) *Writer {
	return &Writer{
		writer: wt,
	}
}

// ----------------------------------------------------------------------------
