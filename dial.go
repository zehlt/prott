package prott

import (
	"fmt"
	"net"
)

func Dial(t ProtocolType, address string) (Conn, error) {
	switch t {

	case TCP:
		conn, err := net.Dial("tcp", address)
		if err != nil {
			return nil, err
		}
		return newTcpConn(conn), nil

	case UDP:
		panic("udp not impl")

	default:
		return nil, fmt.Errorf("protocol not supported")
	}
}
