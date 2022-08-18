package prott

import (
	"net"
)

func NewTcpListener(port string) (Listener, error) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	return &tcpListener{l: l}, nil
}

type tcpListener struct {
	l net.Listener
}

func (w *tcpListener) Accept() (Connection, error) {
	conn, err := w.l.Accept()
	if err != nil {
		return nil, err
	}

	// // // TODO: Use UUID or entity arena
	return newGoConnection(conn), nil
}

func (w *tcpListener) Close() error {
	return w.l.Close()
}

func (w *tcpListener) Addr() string {
	return w.l.Addr().String()
}
