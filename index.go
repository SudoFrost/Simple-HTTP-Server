package main

import "fmt"

func main() {
	server := NewServer(":8080", func(req *Request, res *Response) {
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
	})
	server.AcceptLoop()
}
