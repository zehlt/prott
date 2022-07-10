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

func (r *router) waitForConnection(ctx context.Context, wait Waiter) chan Status {
	statusChan := make(chan Status)

	go func() {
		defer close(statusChan)

		for {
			// TODO: handle when the connection broke or something
			conn, err := wait.Accept()
			if err != nil {
				panic(err)
			}

			statusChan <- Status{T: CONNECTION_STATUS, Connection: conn}
		}

	}()

	return statusChan
}

func (r *router) route(ctx context.Context, statusChan chan Status, messageChan chan Message) {

	go func() {
		// TODO: need to remove chan when close after deconnection from player
		connections := make(map[int]chan network.Packet)

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop

			case status := <-statusChan:
				switch status.T {
				case CONNECTION_STATUS:
					sendChannel := make(chan network.Packet)
					connections[status.Connection.Id()] = sendChannel
					// TODO: maybe move this to his own goroutine
					r.handleClientConnection(ctx, status.Connection, sendChannel, messageChan, statusChan)

				case DISCONNECTION_STATUS:
					send, ok := connections[status.Connection.Id()]
					if !ok {
						panic("try do disconnect no connection")
					}
					close(send)

					delete(connections, status.Connection.Id())
				}

			case message := <-messageChan:
				switch message.T {

				case UNICAST_MESSAGE:
					send, ok := connections[message.Receiver]
					if !ok {
						panic("try to send message to unknown connection router.go")
					}
					send <- message.P

				case BROADCAST_MESSAGE:
					fmt.Println("Broadcast connections", connections)

					for key, send := range connections {
						if key == message.Sender {
							fmt.Println("ingore", key)
							continue
						}

						fmt.Println("send to", key, message, send)
						send <- message.P
					}

				case EMIT_MESSAGE:
					for _, send := range connections {
						send <- message.P
					}
				}
			}
		}
	}()

}

func (r *router) handleClientConnection(ctx context.Context, conn network.Connection, send <-chan network.Packet, messageChan chan<- Message, statusChan chan<- Status) {
	go func() {
		defer conn.Close()

		// TODO: send a deconection packet to signal to not send to him anymore
		ctx, cancel := context.WithCancel(ctx)

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

			statusChan <- Status{T: DISCONNECTION_STATUS, Connection: conn}
		}()

		// TODO: make the dispatch in separate goroutine
		// CONNECTION
		r.dispatch.Disp(network.USER_CONNECTED_PACKET,
			Env{Req: Request{
				Id:     conn.Id(),
				Addr:   conn.RemoteAddr(),
				Packet: network.Packet{T: network.USER_CONNECTED_PACKET, Data: network.UserConnectedPacket{}},
			}, Res: Response{id: conn.Id(), messageChan: messageChan}})

		// receive
		for {
			packet, err := conn.Read()
			if err != nil {
				cancel()

				r.dispatch.Disp(network.USER_DISCONNECTED_PACKET,
					Env{Req: Request{
						Id:     conn.Id(),
						Addr:   conn.RemoteAddr(),
						Packet: network.Packet{T: network.USER_DISCONNECTED_PACKET, Data: network.UserConnectedPacket{}},
					}, Res: Response{}})

				break
			}

			r.dispatch.Disp(packet.T,
				Env{Req: Request{
					Id:     conn.Id(),
					Addr:   conn.RemoteAddr(),
					Packet: packet,
				}, Res: Response{id: conn.Id(), messageChan: messageChan}})
		}
	}()
}
