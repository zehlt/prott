package prott

type Bus struct {
	r          <-chan Packet
	s          chan SocketMessage
	dispatcher Dispatch[PacketType, Packet]
}

func (b *Bus) Init(r <-chan Packet) <-chan SocketMessage {
	var dispatcher Dispatch[PacketType, Packet]
	dispatcher.Init()

	b.r = r
	b.dispatcher = dispatcher
	b.s = make(chan SocketMessage, 1)

	return b.s
}

func (b *Bus) Register(t PacketType, handler func(p Packet)) {
	b.dispatcher.Register(t, handler)
}

func (b *Bus) Send(p Packet) {
	b.s <- SocketMessage{
		T: SEND_MESSAGE_SOCKET,
		P: p,
	}
}

func (b *Bus) Recv() {
	if len(b.r) > 0 {
		p, ok := <-b.r
		if ok {
			b.dispatcher.Disp(p.T, p)
		}
	}
}

func (b *Bus) Close() {
	b.s <- SocketMessage{
		T: CLOSE_MESSAGE_SOCKET,
		P: Packet{},
	}
	close(b.s)
}
