package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server started on port 8080")

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("Client connected")

	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, world!"))
}
