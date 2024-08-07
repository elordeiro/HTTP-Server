package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"net"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Constructor ----------------------------------------------------------------
func NewServer(config *Config) *Server {
	l, err := net.Listen("tcp", "0.0.0.0:"+config.Port)
	if err != nil {
		fmt.Println("Failed to bind to port " + config.Port)
		os.Exit(1)
	}

	return &Server{
		Listener:  l,
		Paths:     map[string]func(*Request, *Server) *Response{},
		Directory: config.Directory,
		Port:      config.Port,
		Encodings: []string{"gzip"},
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
		server.preliminaryChecks(request, response)
		wt.Write(response)
	}
}

func (server *Server) AddPath(path string, handler func(*Request, *Server) *Response) {
	server.Paths[path] = handler
}

func (server *Server) AddEncoding(encoding string) {
	server.Encodings = append(server.Encodings, encoding)
}

func (server *Server) RemoveEncoding(encoding string) {
	for i, e := range server.Encodings {
		if e == encoding {
			server.Encodings = append(server.Encodings[:i], server.Encodings[i+1:]...)
			break
		}
	}
}

func (s *Server) preliminaryChecks(request *Request, response *Response) {
	if e, ok := request.Headers["accept-encoding"]; ok {
		encodings := strings.Split(e, ",")
		for _, encoding := range encodings {
			if slices.Contains(s.Encodings, strings.Trim(encoding, " ")) {
				response.Headers["Content-Encoding"] = encoding
				var buf strings.Builder
				gz := gzip.NewWriter(&buf)
				gz.Write([]byte(response.Body))
				gz.Close()
				response.Body = buf.String()
				response.Headers["Content-Length"] = strconv.Itoa(len(response.Body))
				break
			}
		}
	}
}

func parseFlags() (*Config, error) {
	config := &Config{}
	flag.StringVar(&config.Directory, "directory", "/static/", "Directory that server can access")
	flag.StringVar(&config.Port, "port", "4221", "Server Listening Port")

	flag.Parse()
	return config, nil
}

func main() {
	config, err := parseFlags()
	if err != nil {
		fmt.Println("Failed to parse flags: ", err.Error())
		os.Exit(1)
	}

	server := NewServer(config)

	server.AddPath("/", (*Request).Ok)
	server.AddPath("/echo", (*Request).Echo)
	server.AddPath("/user-agent", (*Request).UserAgent)
	server.AddPath("/files", (*Request).Files)

	server.Listen()
}
