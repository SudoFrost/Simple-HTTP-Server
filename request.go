package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Header map[string][]string

func NewHeader() Header {
	return make(Header)
}

func (h *Header) Set(key string, value string) {
	header := *h
	header[key] = []string{value}
}

func (header Header) Has(key string) bool {
	if header[key] == nil {
		return false
	}
	if len(header[key]) == 0 {
		return false
	}
	return true
}

func (h *Header) Add(key string, value string) {
	header := *h
	if !header.Has(key) {
		header[key] = []string{}
	}
	header[key] = append(header[key], value)
}

func (header Header) Get(key string) []string {
	return header[key]
}

func (h *Header) Delete(key string) {
	header := *h
	if header.Has(key) {
		delete(header, key)
	}
}

func CreateHeader(reader *bufio.Reader) (header Header, err error) {
	header = NewHeader()

	var line string
	for {
		line, err = ReadLine(reader)
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			return
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			err = fmt.Errorf("invalid header line")
			return
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		header.Add(key, value)
	}
}

type Request struct {
	Method string
	Path   string
	Header Header
	Body   []byte
}

func ReadLine(reader *bufio.Reader) (line string, err error) {
	line, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.Trim(line, "\r")
	return
}

func CreateRequest(conn net.Conn) (req *Request) {
	req = &Request{}
	req.Header = make(Header)
	reader := bufio.NewReader(conn)
	line, err := ReadLine(reader)
	if err != nil {
		panic(err)
	}
	parts := strings.Fields(string(line))

	if len(parts) != 3 {
		panic("invalid request line")
	}

	req.Method = parts[0]
	req.Path = parts[1]
	headers, err := CreateHeader(reader)

	if err != nil {
		panic(err)
	}
	req.Header = headers
	return
}
