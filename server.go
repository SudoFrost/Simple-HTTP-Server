package main

import (
	"net"
)

type Server struct {
	listener    net.Listener
	handler     func(req *Request, res *Response)
	connections []net.Conn
}

func NewServer(addr string, handler func(req *Request, res *Response)) *Server {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	server := Server{
		listener:    listener,
		handler:     handler,
		connections: []net.Conn{},
	}

	return &server
}

func (s *Server) CloseConnection(c *net.Conn) {
	conn := *c
	conn.Close()
	tmp := s.connections
	s.connections = []net.Conn{}
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == conn {
			continue
		}
		s.connections = append(s.connections, tmp[i])
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer s.CloseConnection(&conn)
	req := CreateRequest(conn)
	res := NewResponse()

	s.handler(req, res)

	WriteResponse(conn, res)

}

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}
		s.connections = append(s.connections, conn)
		go s.handleConnection(conn)
	}
}
