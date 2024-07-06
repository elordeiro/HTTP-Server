package main

import (
	"fmt"
	"net"
	"os"
)

type Response struct {
	Status  int
	Headers []string
	Message string
}

func OkResponse() []byte {
	return []byte("HTTP/1.1 200 OK\r\n\r\n")
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	conn.Write(OkResponse())
}
