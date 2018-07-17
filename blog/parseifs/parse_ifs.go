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
	requestText, err := readCRLFLine(reader) //string, err Response
	if err != nil {
		//No input, or it doesn't end in CRLF
		return nil, err
	}

	requestLine, err := parseRequestLine(requestText) //RequestLine, err Response
	if err != nil {
		//Not a well-formed HTTP request line with {method, target, version}
		return nil, err
	}

	if request := router.routeRequest(requestLine); request != nil {
		//Well-formed, executable Request to a known route
		return request, nil
	}

	//Valid request, but no route to handle it
	return nil, requestLine.NotImplemented()
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
