package server

import (
	"net"

	"github.com/zehlt/prott/network"
)

func NewTcpWaiter(port string) (Waiter, error) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	return &tcpWaiter{l: l}, nil
}

type tcpWaiter struct {
	l   net.Listener
	ids int
}

func (w *tcpWaiter) Accept() (network.Connection, error) {
	conn, err := w.l.Accept()
	if err != nil {
		return nil, err
	}

	// TODO: Use UUID or entity arena
	w.ids++
	return network.NewGoConnection(conn, w.ids), nil
}

func (w *tcpWaiter) Close() error {
	return w.l.Close()
}

func (w *tcpWaiter) Addr() string {
	return w.l.Addr().String()
}
