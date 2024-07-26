package main

import (
	"fmt"
	"net"
	"os"
)

// Constructor ----------------------------------------------------------------
func NewServer(port string) *Server {
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println("Failed to bind to port " + port)
		os.Exit(1)
	}

	paths := make([]string, 0)
	paths = append(paths, "/")

	return &Server{
		Listener: l,
		Paths:    paths,
	}
}

// ----------------------------------------------------------------------------

func (server *Server) Listen() {
	for {
		conn, err := server.Listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go server.Serve(conn)
	}
}

func (server *Server) Serve(conn net.Conn) {
	rd := NewReader(conn)
	wt := NewWriter(conn)
	for {
		request, err := rd.Read()
		if err != nil {
			fmt.Println("Error reading request: ", err.Error())
			continue
		}
		response := server.Handle(request)
		wt.Write(response)
	}
}

func OkResponse() []byte {
	return []byte("HTTP/1.1 200 OK\r\n\r\n")
}

func main() {
	server := NewServer("4221")
	server.Listen()
}
