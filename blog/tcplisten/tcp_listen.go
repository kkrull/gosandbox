package tcplisten

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := StartListening("localhost", 1)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("Listening on: %s", listener.Addr())
}

func StartListening(host string, port uint16) (net.Listener, error) {
	address, addressErr := net.ResolveTCPAddr("blog", fmt.Sprintf("%s:%d", host, port))
	if addressErr != nil {
		return nil, fmt.Errorf("StartListening: %s", addressErr)
	}

	listener, listenError := net.ListenTCP("blog", address)
	if listenError != nil {
		return nil, fmt.Errorf("StartListening: %s", listenError)
	}

	return listener, nil
}
