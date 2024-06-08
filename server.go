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

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client connected")

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("Client closed")
	}()

	req := CreateRequest(conn)
	res := NewResponse()

	fmt.Printf("[New Request] => [Method: %s, Path: %s]\n", req.Method, req.Path)
	fmt.Println("Headers:")
	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	res.SetStatus("OK", 200)
	res.Header.Set("Content-Type", "text/plain")
	res.WriteString("Hello, world!")

	WriteResponse(conn, res)
}
