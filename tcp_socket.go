package prott

import (
	"context"
	"fmt"
	"net"
)

type tcpSocket struct {
	isConnected bool
	cancel      context.CancelFunc
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

	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel

	bus := handleServerConnection(ctx, newGoConnection(conn, 0))
	return bus, nil
}

func (c *tcpSocket) Disconnect() {
	c.cancel()
	c.isConnected = false
}

func handleServerConnection(ctx context.Context, conn Connection) Bus {
	recv := make(chan Packet, 1)
	send := make(chan Packet, 1)

	go func() {
		cont, cancel := context.WithCancel(ctx)

		// send
		go func() {
			defer conn.Close()
			defer close(recv)
			defer close(send)

		out:
			for {
				select {
				case packet_to_send := <-send:
					err := conn.Write(packet_to_send)
					if err != nil {
						panic(err)
					}

				case <-cont.Done():
					break out
				}
			}
		}()

		// receive
		for {
			packet, err := conn.Read()
			if err != nil {
				cancel()
				break
			}

			recv <- packet
		}
	}()

	return Bus{recvC: recv, sendC: send}
}
