package prott

import (
	"context"
	"fmt"
	"net"
)

type tcpSocket struct {
	isConnected bool
}

func newTcpSocket() Socket {
	return &tcpSocket{
		isConnected: false,
	}
}

func (c *tcpSocket) Connect(port string) (Bus, error) {
	if c.isConnected {
		return Bus{}, fmt.Errorf("already connected")
	}

	conn, err := net.Dial("tcp", port)
	if err != nil {
		return Bus{}, err
	}
	c.isConnected = true

	bus := c.handleServerConnection(newGoConnection(conn, 0))
	return bus, nil
}

func (c *tcpSocket) handleServerConnection(conn Connection) Bus {
	recv := make(chan Packet, 1)

	var bus Bus
	send := bus.Init(recv)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer conn.Close()

	loop:
		for {
			select {

			case <-ctx.Done():
				break loop

			case socketMessage := <-send:
				switch socketMessage.T {

				case SEND_MESSAGE_SOCKET:
					err := conn.Write(socketMessage.P)
					if err != nil {
						panic(fmt.Sprintln("error trying to send packet", err))
					}

				case CLOSE_MESSAGE_SOCKET:
					break loop
				}
			}
		}
		c.isConnected = false
	}()

	go func() {
		defer close(recv)

		for {
			packet, err := conn.Read()
			if err != nil {
				recv <- Packet{T: SERVER_CLOSE_CONNECTION_PACKET, Data: ServerCloseConnectionPacket{Err: err.Error()}}
				cancel()
				break
			}
			recv <- packet
		}
	}()

	return bus
}
