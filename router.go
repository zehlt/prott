package prott

import (
	"context"
)

type Router interface {
	Register(t PacketType, f func(Env))
	Serve(ctx context.Context, wait Waiter) chan<- Message
}

type router struct {
	dispatch Dispatch[PacketType, Env]
}

func newRouter() Router {
	r := router{}
	r.dispatch.Init()

	return &r
}

func (r *router) Register(t PacketType, f func(Env)) {
	r.dispatch.Register(t, f)
}

func (r *router) Serve(ctx context.Context, wait Waiter) chan<- Message {
	// TODO: close messageChan
	// defer close(messageChan)
	messageChan := make(chan Message)

	connections := r.waitForConnection(ctx, wait)
	r.route(ctx, connections, messageChan)

	return messageChan
}

func (r *router) waitForConnection(ctx context.Context, wait Waiter) chan status {
	statusChan := make(chan status)

	go func() {
		defer close(statusChan)

		for {
			// TODO: handle when the connection broke or something
			conn, err := wait.Accept()
			if err != nil {
				panic(err)
			}

			statusChan <- status{t: CONNECTION_STATUS, connection: conn}
		}

	}()

	return statusChan
}

func (r *router) route(ctx context.Context, statusChan chan status, messageChan chan Message) {

	go func() {
		// TODO: need to remove chan when close after deconnection from player
		connections := make(map[int]chan Packet)

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop

			case status := <-statusChan:
				switch status.t {
				case CONNECTION_STATUS:
					sendChannel := make(chan Packet)
					connections[status.connection.Id()] = sendChannel
					// TODO: maybe move this to his own goroutine
					r.handleClientConnection(ctx, status.connection, sendChannel, messageChan, statusChan)

				case DISCONNECTION_STATUS:
					send, ok := connections[status.connection.Id()]
					if !ok {
						panic("try do disconnect no connection")
					}
					close(send)
					delete(connections, status.connection.Id())
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

func (r *router) handleClientConnection(ctx context.Context, conn Connection, send <-chan Packet, messageChan chan<- Message, statusChan chan<- status) {

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

		statusChan <- status{t: DISCONNECTION_STATUS, connection: conn}
	}()

	go func() {
		defer conn.Close()

		// TODO: make the dispatch in separate goroutine
		// CONNECTION
		r.dispatch.Disp(USER_CONNECTED_PACKET,
			Env{Req: Request{
				id:     conn.Id(),
				addr:   conn.RemoteAddr(),
				packet: Packet{T: USER_CONNECTED_PACKET, Data: UserConnectedPacket{}},
			}, Res: Response{id: conn.Id(), messageChan: messageChan}})

		// receive
		for {
			packet, err := conn.Read()
			if err != nil {
				cancel()

				r.dispatch.Disp(USER_DISCONNECTED_PACKET,
					Env{Req: Request{
						id:     conn.Id(),
						addr:   conn.RemoteAddr(),
						packet: Packet{T: USER_DISCONNECTED_PACKET, Data: UserConnectedPacket{}},
					}, Res: Response{}})

				break
			}

			r.dispatch.Disp(packet.T,
				Env{Req: Request{
					id:     conn.Id(),
					addr:   conn.RemoteAddr(),
					packet: packet,
				}, Res: Response{id: conn.Id(), messageChan: messageChan}})
		}
	}()
}
