package prott

import "net"

type tcpSocket struct {
}

func NewTcpSocket() Socket {
	return &tcpSocket{}
}

func (s *tcpSocket) Connect(address string) (Connection, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return newGoConnection(conn), nil
}
