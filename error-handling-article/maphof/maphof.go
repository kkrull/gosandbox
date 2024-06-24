package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kkrull/gohttp/msg/clienterror"
	"github.com/kkrull/gohttp/msg/servererror"
)

func main() {
	router := &HttpRouter{}
	router.AddRoute(&getRoute{})

	request, err := router.parse(bufio.NewReader(bytes.NewBufferString("GET / HTTP/1.1\r\n\r\n")))
	if err != nil {
		fmt.Printf("parse_ifs: %s", err)
		os.Exit(1)
	}

	request.Handle(os.Stdout)
}

type HttpRouter struct {
	routes []Route
}

func (router *HttpRouter) AddRoute(route Route) {
	router.routes = append(router.routes, route)
}

func (router HttpRouter) parse(reader *bufio.Reader) (Request, Response) {
	request, response := newStringOrResponse(readCRLFLine(reader)).
		Map(parseRequestLine).
		Map(router.routeRequest)

	if response != nil {
		return nil, response
	} else if request == nil {
		// Technically, this doesn't work because we now lack the intermediate value
		return nil, &servererror.NotImplemented{}
	} else {
		return request, nil
	}
}

func readCRLFLine(reader *bufio.Reader) (string, Response) {
	maybeEndsInCR, _ := reader.ReadString('\r')
	if len(maybeEndsInCR) == 0 {
		return "", &clienterror.BadRequest{DisplayText: "end of input before terminating CRLF"}
	} else if !strings.HasSuffix(maybeEndsInCR, "\r") {
		return "", &clienterror.BadRequest{DisplayText: "line in request header not ending in CRLF"}
	}

	nextCharacter, _ := reader.ReadByte()
	if nextCharacter != '\n' {
		return "", &clienterror.BadRequest{DisplayText: "message header line does not end in LF"}
	}

	trimmed := strings.TrimSuffix(maybeEndsInCR, "\r")
	return trimmed, nil
}

func newStringOrResponse(data string, err Response) *StringOrResponse {
	return &StringOrResponse{data: data, err: err}
}

type StringOrResponse struct {
	data string
	err  Response
}

type ParseRequestLine func(text string) (*RequestLine, Response)

func (either *StringOrResponse) Map(parse ParseRequestLine) *RequestLineOrResponse {
	if either.err != nil {
		return &RequestLineOrResponse{data: nil, err: either.err}
	}

	requestLine, err := parse(either.data)
	if err != nil {
		return &RequestLineOrResponse{data: nil, err: either.err}
	}

	return &RequestLineOrResponse{data: requestLine, err: nil}
}

type RequestLineOrResponse struct {
	data *RequestLine
	err  Response
}

type RouteRequest func(requested *RequestLine) Request

func (either *RequestLineOrResponse) Map(route RouteRequest) (Request, Response) {
	if either.err != nil {
		return nil, either.err
	}

	return route(either.data), nil
}

func parseRequestLine(text string) (*RequestLine, Response) {
	fields := strings.Split(text, " ")
	if len(fields) != 3 {
		return nil, &clienterror.BadRequest{DisplayText: "incorrectly formatted or missing request-line"}
	}

	return &RequestLine{
		Method: fields[0],
		Target: fields[1],
	}, nil
}

func (router HttpRouter) routeRequest(requested *RequestLine) Request {
	for _, route := range router.routes {
		request := route.Route(requested)
		if request != nil {
			return request
		}
	}

	return nil
}

type getRoute struct{}

func (*getRoute) Route(requested *RequestLine) Request {
	return &getRequest{}
}

type getRequest struct{}

func (request *getRequest) Handle(client io.Writer) error {
	fmt.Fprintf(client, "HTTP/1.1 204 No Content\r\n\r\n")
	return nil
}

type Route interface {
	Route(requested *RequestLine) Request
}

type RequestLine struct {
	Method string
	Target string
}

func (requestLine *RequestLine) NotImplemented() Response {
	return &servererror.NotImplemented{Method: requestLine.Method}
}

type Request interface {
	Handle(client io.Writer) error
}

type Response interface {
	WriteTo(client io.Writer) error
	WriteHeader(client io.Writer) error
}
