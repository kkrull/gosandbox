package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := Listen("localhost", 54321)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("Listening on: %s", listener.Addr())
}

func Listen(host string, port uint16) (net.Listener, error) {
	addr, addrErr := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
	if addrErr != nil {
		return nil, fmt.Errorf("Listen: %s", addrErr)
	}

	listener, listenError := net.ListenTCP("tcp", addr)
	if listenError != nil {
		return nil, fmt.Errorf("Listen: %s", listenError)
	}

	return listener, nil
}
