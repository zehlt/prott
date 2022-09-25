package prott

import (
	"fmt"
	"net"
)

type Listener interface {
	Accept() (Conn, error)
	Close() error
	// Addr() Addr
}

func Listen(t ProtocolType, address string) (Listener, error) {
	switch t {
	case TCP:
		l, err := net.Listen("tcp", address)
		if err != nil {
			return nil, err
		}
		return &tcpListener{l: l}, nil
	case UDP:
		panic("udp not impl")
	default:
		return nil, fmt.Errorf("protocol not supported")
	}
}

type tcpListener struct {
	l net.Listener
}

func (tl *tcpListener) Accept() (Conn, error) {
	conn, err := tl.l.Accept()
	if err != nil {
		return nil, err
	}

	return newTcpConn(conn), nil
}

func (tl *tcpListener) Close() error {
	return tl.l.Close()
}
