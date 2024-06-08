package main

import (
	"fmt"
	"net"
	"strconv"
)

type Response struct {
	Status string
	Code   int
	Header Header
	Body   []byte
}

func NewResponse() *Response {
	return &Response{
		Status: "200 OK",
		Code:   200,
		Header: NewHeader(),
		Body:   []byte{},
	}
}

func (r *Response) SetStatus(status string, code int) {
	r.Status = status
	r.Code = code
}

func (r *Response) write(data []byte) {
	r.Body = data
	r.Header.Set("Content-Length", strconv.Itoa(len(data)))
}

func (r *Response) WriteString(data string) {
	r.write([]byte(data))
}

func WriteResponse(writer net.Conn, res *Response) {
	writer.Write([]byte(fmt.Sprintf("HTTP/1.1 %d %s\r\n", res.Code, res.Status)))
	for k, values := range res.Header {
		for _, v := range values {
			writer.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
		}
	}
	writer.Write([]byte("\r\n"))
	writer.Write(res.Body)
}
