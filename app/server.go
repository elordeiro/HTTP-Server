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

	return &Server{
		Listener: l,
		Paths:    map[string]func(*Request) *Response{},
	}
}

// ----------------------------------------------------------------------------

func (server *Server) Listen() {
	fmt.Println("Listening on port " + server.Listener.Addr().String())
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
			if err.Error() == "EOF" {
				fmt.Println("Connection closed")
				conn.Close()
				break
			}
			fmt.Println("Error reading request: ", err.Error())
			return
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
	server.AddPath("/", (*Request).Ok)
	server.AddPath("/echo", (*Request).Echo)
	server.Listen()
}
