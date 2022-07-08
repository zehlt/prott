package server

import (
	"context"
	"fmt"
	"time"

	"github.com/zehlt/prott/network"
)

type Router interface {
	Register(t network.PacketType, f func(Env))
	Serve(ctx context.Context, wait Waiter)
}

type router struct {
	dispatch network.Dispatch[network.PacketType, Env]
}

func NewRouter() Router {
	r := router{}
	r.dispatch.Init()

	return &r
}

func (r *router) Register(t network.PacketType, f func(Env)) {
	r.dispatch.Register(t, f)
}

func (r *router) Serve(ctx context.Context, wait Waiter) {
	messageChan := make(chan Message)
	defer close(messageChan)

	connections := r.waitForConnection(ctx, wait)
	r.route(ctx, connections, messageChan)

	time.Sleep(time.Hour * 1)
}

func (r *router) waitForConnection(ctx context.Context, wait Waiter) <-chan network.Connection {
	connChan := make(chan network.Connection)

	go func() {
		defer close(connChan)
		fmt.Println("waiting for connection...")

		for {
			// TODO: handle when the connection broke or something
			conn, err := wait.Accept()
			if err != nil {
				panic(err)
			}

			connChan <- conn
		}

	}()

	return connChan
}

func (r *router) route(ctx context.Context, connChan <-chan network.Connection, messageChan chan Message) {

	go func() {
		// TODO: need to remove chan when close after deconnection from player
		connections := make(map[int]chan network.Packet)

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop

			case conn := <-connChan:
				sendChannel := make(chan network.Packet)
				connections[conn.Id()] = sendChannel
				r.handleClientConnection(ctx, conn, sendChannel, messageChan)

			case message := <-messageChan:
				switch message.T {

				case UNICAST_MESSAGE:
					fmt.Println("SEND UNICAST")
					send, ok := connections[message.Receiver]
					if !ok {
						panic("try to send message to unknown connection router.go")
					}
					send <- message.P

				case BROADCAST_MESSAGE:
					fmt.Println("SEND BROADCAST")
					for key, send := range connections {
						fmt.Println(key)
						fmt.Println(send)
					}

				case EMIT_MESSAGE:
					fmt.Println("SEND EMIT")

					for _, send := range connections {
						send <- message.P
					}
				}
			}
		}
	}()

}

func (r *router) handleClientConnection(ctx context.Context, conn network.Connection, send <-chan network.Packet, messageChan chan<- Message) {
	go func() {
		defer conn.Close()

		// TODO: send a deconection packet to signal to not send to him anymore
		ctx, cancel := context.WithCancel(ctx)

		// Handshake
		err := conn.Write(network.Packet{T: network.SERVER_CONNECTION_ACCEPTED_PACKET, Data: network.ServerConnectionAcceptedPacket{}})
		if err != nil {
			panic(err)
		}

		// CONNECTION
		r.dispatch.Disp(network.USER_CONNECTED_PACKET,
			Env{Req: Request{
				Id:     conn.Id(),
				Addr:   conn.RemoteAddr(),
				Packet: network.Packet{T: network.USER_CONNECTED_PACKET, Data: network.UserConnectedPacket{}},
			}, Res: Response{id: conn.Id(), messageChan: messageChan}})

		// send
		go func() {
		out:
			for {
				select {
				case packet_to_send := <-send:
					err := conn.Write(packet_to_send)
					if err != nil {
						panic(err)
					}

				case <-ctx.Done():
					break out
				}
			}

			// DISCONNECTION
			r.dispatch.Disp(network.USER_DISCONNECTED_PACKET,
				Env{Req: Request{
					Id:     conn.Id(),
					Addr:   conn.RemoteAddr(),
					Packet: network.Packet{T: network.USER_DISCONNECTED_PACKET, Data: network.UserConnectedPacket{}},
				}, Res: Response{}})
		}()

		// receive
		for {
			packet, err := conn.Read()
			if err != nil {
				cancel()
				break
			}

			r.dispatch.Disp(packet.T,
				Env{Req: Request{
					Id:     conn.Id(),
					Addr:   conn.RemoteAddr(),
					Packet: packet,
				}, Res: Response{}})
		}
	}()
}
