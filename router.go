package prott

import (
	"context"
	"time"
)

type Router interface {
	Register(t PacketType, f func(Env))
	Serve(ctx context.Context, wait Waiter)
}

type router struct {
	dispatch Dispatch[PacketType, Env]
}

func NewRouter() Router {
	r := router{}
	r.dispatch.Init()

	return &r
}

func (r *router) Register(t PacketType, f func(Env)) {
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
		connections := make(map[int]chan Packet)

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop

			case status := <-statusChan:
				switch status.T {
				case CONNECTION_STATUS:
					sendChannel := make(chan Packet)
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
					for key, send := range connections {
						if key == message.Sender {
							continue
						}

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

func (r *router) handleClientConnection(ctx context.Context, conn Connection, send <-chan Packet, messageChan chan<- Message, statusChan chan<- Status) {
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
		r.dispatch.Disp(USER_CONNECTED_PACKET,
			Env{Req: Request{
				Id:     conn.Id(),
				Addr:   conn.RemoteAddr(),
				Packet: Packet{T: USER_CONNECTED_PACKET, Data: UserConnectedPacket{}},
			}, Res: Response{id: conn.Id(), messageChan: messageChan}})

		// receive
		for {
			packet, err := conn.Read()
			if err != nil {
				cancel()

				r.dispatch.Disp(USER_DISCONNECTED_PACKET,
					Env{Req: Request{
						Id:     conn.Id(),
						Addr:   conn.RemoteAddr(),
						Packet: Packet{T: USER_DISCONNECTED_PACKET, Data: UserConnectedPacket{}},
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
